package Requests

import (
	"context"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"reflect"
	"strings"
)

type ValidationRule struct {
	Rules        map[string]interface{}
	ErrorMessage map[string]interface{}
}

var v = validator.New()

func validate(ctx context.Context, data map[string]interface{}, rules map[string]interface{}) map[string]interface{} {
	errs := make(map[string]interface{})
	for field, rule := range rules {
		if reflect.ValueOf(rule).Kind() == reflect.Map && reflect.ValueOf(data[field]).Kind() == reflect.Map {
			err := validate(ctx, data[field].(map[string]interface{}), rule.(map[string]interface{}))
			if len(err) > 0 {
				errs[field] = err
			}
		} else if reflect.ValueOf(rule).Kind() == reflect.Map {
			errs[field] = errors.New("The field: '" + field + "' is not a map to dive")
		} else {
			err := v.VarCtx(ctx, data[field], rule.(string))
			if err != nil {
				errs[field] = err
			}
		}
	}
	return errs
}

func ValidateForm(ctx iris.Context, validationRules ValidationRule) map[string]interface{} {
	data := make(map[string]interface{})
	for field, value := range ctx.FormValues() {
		data[field] = strings.Join(value, "")
	}
	return parseErrors(validate(context.Background(), data, validationRules.Rules),validationRules.ErrorMessage)
}

func parseErrors(errors map[string]interface{}, errorList map[string]interface{}) (errs map[string]interface{}) {
	errs = make(map[string]interface{})
	for field, err := range errors {
		switch reflect.ValueOf(err).Kind() == reflect.Map {
		case true:
			errs[field] = parseErrors(errors[field].(map[string]interface{}), errorList[field].(map[string]interface{}))
			break
		default:
			errs[field] = errorList[field]
		}
	}
	return
}
