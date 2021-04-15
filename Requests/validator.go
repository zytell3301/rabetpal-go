package Requests

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"strings"
)

type ValidationRule struct {
	Rules        map[string]string
	ErrorMessage map[string]string
}

var validatorPackage = validator.New()

func Validate(data map[string]interface{}, validationRules ValidationRule) (bool, map[string]string) {
	rules := validationRules.Rules
	validationErrors := make(map[string]string)
	var err error = nil
	for field, rule := range rules {
		value := data[field]
		err = validatorPackage.Var(value, rule)
		switch err != nil {
		case true:
			validationErrors[field] = validationRules.ErrorMessage[field]
			err = nil
		}
	}
	return true, validationErrors
}

func ValidateForm(ctx iris.Context, validationRules ValidationRule) (bool, map[string]string) {
	data := make(map[string]interface{})
	for field, value := range ctx.FormValues() {
		data[field] = strings.Join(value, "")
	}
	return Validate(data, validationRules)
}
