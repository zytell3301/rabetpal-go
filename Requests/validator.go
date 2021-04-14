package Requests

import (
	"github.com/go-playground/validator/v10"
)

var validatorPackage = validator.New()

func Validate(data map[string]interface{}, rules map[string]string) bool {
	for offset, rule := range rules {
		value, isSet := data[offset]
		switch isSet {
		case true:
			validatorPackage.Var(value, rule)
			break
		default:
			validatorPackage.Var(nil, rule)
		}
	}
	return true
}
