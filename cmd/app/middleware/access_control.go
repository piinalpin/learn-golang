package middleware

import (
	"learn-rest-api/cmd/app/component"
	"learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/exception"
	"learn-rest-api/pkg"
	"strings"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RoleBasedAccessControl(hasRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, isExists := c.Get(constant.JwtClaims.GetContextKey())
		if !isExists {
			log.Error("Access Denied: JWT claims not found in context.")
			exception.ThrowNewAppException(constant.AccessDenied)
		}

		roles := strings.Split(claims.(*component.JwtCustomClaims).Roles, ",")

		if !pkg.Contains(roles, hasRole) {
			log.Error("Access Denied: Don't have access to the resource.")
			exception.ThrowNewAppException(constant.AccessDenied)
		}
		
		c.Next()
	}
}