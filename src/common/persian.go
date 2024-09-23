package common

import (
	"regexp"

	"github.com/paghapour/golang-clean-web-api/config"
	"github.com/paghapour/golang-clean-web-api/pkg/logging"
)

var logger = logging.NewLogger(config.GetConfig())

const iranianMobileNumberPattern string = `^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`

func IranianMobileNumberValidate(mobileNumber string) bool {
	res, err := regexp.MatchString(iranianMobileNumberPattern, mobileNumber)
	if err != nil {
		logger.Error(logging.Postgres, logging.MobileValidation, err.Error(), nil)
	}
	return res
}
