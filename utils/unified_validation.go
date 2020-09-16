package util

import (
	"fmt"
	"reflect"
	"seifwu/global"
	"seifwu/global/response"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UnifiedValidation 统一参数验证
func UnifiedValidation(c *gin.Context, err error, errCode string, params interface{}) {
	errs, ok := err.(validator.ValidationErrors)
	var result = make(map[string]interface{})
	result["success"] = false
	result["errCode"] = errCode

	if !ok {
		// 非 validator.ValidationErrors 类型错误直接返回
		result["errCode"] = err.Error()
		response.Fail(c, result)
		return
	}

	var fieldsTrans = make(map[string]string)
	fields := reflect.TypeOf(params)
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		keyName := field.Tag.Get("json")
		value := field.Tag.Get("label")
		fieldsTrans[keyName] = value
	}

	var resultErrs = make(map[string]string)
	for k, v := range ValidatorCustomErrorName(errs.Translate(global.Trans)) {
		fmt.Println(k, v)
		v = strings.Replace(v, k, fieldsTrans[k], -1)
		resultErrs[k] = v
	}

	result["errMsg"] = resultErrs
	response.Fail(c, result)
	return
}
