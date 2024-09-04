package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/paghapour/golang-clean-web-api/common"
)

// ^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$
func IranianMobileNumberValidator(fld validator.FieldLevel) bool {
	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}


	return common.IranianMobileNumberValidate(value)
}
