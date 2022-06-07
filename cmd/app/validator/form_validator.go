package validator

import (
	"errors"
	"fmt"
	respkey "learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/exception"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindJSON(c *gin.Context, form any) {
	var err = c.ShouldBindJSON(form)
	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			exception.ThrowNewAppException_(respkey.InvalidRequest.GetKey(), getErrorMsg(ve[0]))
		}
	}

}

func getErrorMsg(fe validator.FieldError) string {
    switch fe.Tag() {
        case "required":
            return fmt.Sprintf("`%s` is required", fe.Field())
        case "lte":
            return fmt.Sprintf("Field `%s` must be less than or equal to %s", fe.Field(), fe.Param())
        case "gte":
			return fmt.Sprintf("Field `%s` should be greater than %s", fe.Field(), fe.Param())
		case "min":
			return fmt.Sprintf("Field `%s` should be at least %s characters", strings.ToLower(fe.Field()), fe.Param())
		case "max":
			return fmt.Sprintf("Field `%s` should be at most %s characters", strings.ToLower(fe.Field()), fe.Param())
		default:
			return "Invalid field"
    }
}	