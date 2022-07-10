package middleware

import (
	"learn-rest-api/cmd/app/component"
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/exception"
	"strings"

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

		t.ValidateAccessToken(token)
		c.Next()
	}
}
