package util

import (
	"fmt"
	"reflect"
	"seifwu/global"
	"strings"

	"github.com/go-playground/validator/v10"
)

// UnifiedValidation 统一参数验证
func UnifiedValidation(err error, params interface{}) interface{} {
	errs, ok := err.(validator.ValidationErrors)

	if !ok {
		// 非 validator.ValidationErrors 类型错误直接返回
		return err.Error()
	}

	var fieldsTrans = make(map[string]string)
	var tagTrans = make(map[string]string)
	fields := reflect.TypeOf(params)
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		keyName := field.Tag.Get("json")
		value := field.Tag.Get("label")
		fieldsTrans[keyName] = value
		tagTrans[field.Name] = value
	}

	var resultErrs = make(map[string]string)
	for k, v := range ValidatorCustomErrorName(errs.Translate(global.Trans)) {
		fmt.Println(k, v)
		v = strings.Replace(v, k, fieldsTrans[k], -1)
		for key, value := range tagTrans {
			v = strings.ReplaceAll(v, key, value)
		}

		resultErrs[k] = v
	}

	return resultErrs
}
