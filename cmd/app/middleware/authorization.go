package middleware

import (
	"learn-rest-api/cmd/app/component"
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/exception"
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Bearer(t component.TokenProvider) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer exception.AppExceptionHandler(c)

		token := c.Request.Header.Get(constant.Authorization.GetHeadersName())
		token = strings.Replace(token, constant.Bearer.GetHeadersPrefix(), "", 1)

		if token == "" {
			exception.ThrowNewAppException(constant.Unauthorized)
		}

		claims := t.ValidateAccessToken(token)
		c.Set(constant.JwtClaims.GetContextKey(), claims)
		c.Next()
	}
}

func Basic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer exception.AppExceptionHandler(c)
		godotenv.Load()

		clientId := os.Getenv("application.client-id")
		clientSecret := os.Getenv("application.client-secret")
		
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			log.Error("Failed parsing basic auth.")
			exception.ThrowNewAppException_(constant.Unauthorized.GetKey(), "Invalid client ID or secret!")
		}

		isValid := (username == clientId) && (password == clientSecret)
		if !isValid {
			log.Error("Invalid credentials!")
			exception.ThrowNewAppException_(constant.Unauthorized.GetKey(), "Invalid client ID or secret!")
		}

		c.Next()
	}
}
