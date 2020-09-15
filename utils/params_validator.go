package util

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// ValidatorCustomErrorName 自定义 error name
func ValidatorCustomErrorName(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func TransTagName(libTans, err interface{}) interface{} {
	switch err.(type) {
	case validator.ValidationErrorsTranslations:
		var errs map[string]string
		errs = make(map[string]string, 0)
		for k, v := range err.(validator.ValidationErrorsTranslations) {
			for key, value := range libTans.(map[string]string) {
				v = strings.Replace(v, key, value, -1)
			}
			errs[k] = v
		}

		return ValidatorCustomErrorName(errs)
	case string:
		var errs string
		for key, value := range libTans.(map[string]string) {
			errs = strings.Replace(errs, key, value, -1)
		}
		return errs
	default:
		return err
	}
}
