package validator

import (
	"errors"
	"fmt"
	"regexp"
	respkey "learn-rest-api/cmd/app/constant"
	"learn-rest-api/cmd/app/exception"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")

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
            return fmt.Sprintf("Field `%s` must be less than or equal to %s", getFieldName(fe.Field()), fe.Param())
        case "gte":
			return fmt.Sprintf("Field `%s` should be greater than %s", getFieldName(fe.Field()), fe.Param())
		case "min":
			return fmt.Sprintf("Field `%s` should be at least %s characters", getFieldName(fe.Field()), fe.Param())
		case "max":
			return fmt.Sprintf("Field `%s` should be at most %s characters", getFieldName(fe.Field()), fe.Param())
		case "len":
			return fmt.Sprintf("Field `%s` must be %s characters", getFieldName(fe.Field()), fe.Param())
		default:
			return "Invalid field validation"
    }
}

func getFieldName(value string) string {
	value = matchFirstCap.ReplaceAllString(value, "${1}_${2}")
    value  = matchAllCap.ReplaceAllString(value, "${1}_${2}")
    return strings.ToLower(value)
}