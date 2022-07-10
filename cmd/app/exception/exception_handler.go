package exception

import (
	"fmt"
	respkey "learn-rest-api/cmd/app/constant"
	"learn-rest-api/pkg"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AppExceptionHandler(c *gin.Context) {
	if err := recover(); err != nil {
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		switch key {
		case
			respkey.DataNotFound.GetKey(),
			respkey.InvalidRequest.GetKey():
			c.JSON(http.StatusBadRequest, pkg.BuildResponse_(key, msg, pkg.Null()))
			c.Abort()
		case
			respkey.Unauthorized.GetKey():
			c.JSON(http.StatusUnauthorized, pkg.BuildResponse_(key, msg, pkg.Null()))
			c.Abort()
		default:
			c.JSON(http.StatusInternalServerError, pkg.BuildResponse_(respkey.UnknownError.GetKey(), msg, pkg.Null()))
			c.Abort()
		}

	}
}
