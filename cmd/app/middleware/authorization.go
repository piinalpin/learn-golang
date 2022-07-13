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
	return func(ctx *gin.Context) {
		defer exception.AppExceptionHandler(ctx)

		token := ctx.Request.Header.Get(constant.Authorization.GetHeadersName())
		token = strings.Replace(token, constant.Bearer.GetHeadersPrefix(), "", 1)

		if token == "" {
			exception.ThrowNewAppException(constant.Unauthorized)
		}

		claims := t.ValidateAccessToken(token)
		ctx.Set(constant.JwtClaims.GetContextKey(), claims)
		ctx.Next()
	}
}

func Basic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer exception.AppExceptionHandler(ctx)
		godotenv.Load()

		clientId := os.Getenv("application.client-id")
		clientSecret := os.Getenv("application.client-secret")
		
		username, password, ok := ctx.Request.BasicAuth()
		if !ok {
			log.Error("Failed parsing basic auth.")
			exception.ThrowNewAppException_(constant.Unauthorized.GetKey(), "Invalid client ID or secret!")
		}

		isValid := (username == clientId) && (password == clientSecret)
		if !isValid {
			log.Error("Invalid credentials!")
			exception.ThrowNewAppException_(constant.Unauthorized.GetKey(), "Invalid client ID or secret!")
		}

		ctx.Next()
	}
}
