package middleware

import (
	"learn-rest-api/cmd/app/component"
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/exception"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Authorization(t component.TokenProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer exception.AppExceptionHandler(c)

		token := c.Request.Header.Get(constant.Authorization.GetHeadersName())
		token = strings.Replace(token, constant.Bearer.GetHeadersPrefix(), "", 1)

		if token == "" {
			exception.ThrowNewAppException(constant.Unauthorized)
		}

		claims := t.ValidateAccessToken(token)
		log.Debug("Claims: ", claims)
		c.Set(constant.JwtClaims.GetContextKey(), claims)
		c.Next()
	}
}
