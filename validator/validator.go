package validator

import (
	"reflect"
	"strings"

	r "github.com/fiuskylab/go-valid/validator/rules"
)

type validator struct {
	Result Result
	values map[string]interface{}
}

func Check(i interface{}, rules r.Rules) (Result, error) {
	typeOf := reflect.TypeOf(i)
	valueOf := reflect.ValueOf(i)

	v := validator{
		Result: make(Result),
		values: make(map[string]interface{}),
	}

	var err error
	for i := 0; i < typeOf.NumField(); i++ {
		fieldNameLower := strings.ToLower(typeOf.Field(i).Name)
		fieldValue := valueOf.Field(i).Interface()
		v.values[fieldNameLower] = fieldValue
	}

	for key, val := range v.values {
		if err = v.validateField(val, key, rules[key]); err != nil {
			return v.Result, err
		}
	}
	return v.Result, err
}

const (
	emptyRulesForField = `no rules found for %s`
)

func (v *validator) validateField(value interface{}, fieldName string, rules r.Rule) error {
	results := []string{}

	var err error

	if rules.Required {
		if res := r.Required(value); res != "" {
			results = append(results, res)
		}
	}

	if rules.Max != 0 {
		if res := r.Max(value, rules.Max); res != "" {
			results = append(results, res)
		}
	}

	if rules.Min != 0 {
		if res := r.Min(value, rules.Min); res != "" {
			results = append(results, res)
		}
	}

	if rules.Regex != "" {
		if res := r.Regex(value, rules.Regex); res != "" {
			results = append(results, res)
		}
	}

	if rules.Type != r.Nil {
		if res := r.IsType(value, rules.Type); res != "" {
			results = append(results, res)
		}
	}

	if !isNil(rules.TypeVar) {
		if res := r.IsTypeVar(value, rules.TypeVar); res != "" {
			results = append(results, res)
		}
	}

	if len(rules.In) > 0 {
		if res := r.In(value, rules.In); res != "" {
			results = append(results, res)
		}
	}

	if len(results) > 0 {
		v.Result[fieldName] = results
	}

	return err
}
