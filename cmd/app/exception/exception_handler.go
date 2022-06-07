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
		fmt.Println(err)
		str := fmt.Sprint(err)
		strArr := strings.Split(str, ":")

		key := strArr[0]
		msg := strings.Trim(strArr[1], " ")

		fmt.Printf("KEY::: %s", key)
		fmt.Printf("MSG::: %s", msg)
		switch key {
			case 
				respkey.UnknownError.GetKey():
				c.JSON(http.StatusInternalServerError, pkg.BuildResponse_(key, msg, pkg.Null()))
			default:
				c.JSON(http.StatusBadRequest, pkg.BuildResponse_(key, msg, pkg.Null()))
		}
		
	}
}