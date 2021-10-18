package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	r "github.com/fiuskylab/go-valid/validator/rules"
)

type validator struct {
	Result Result
}

func Check(i interface{}, rules map[string][]string) (Result, error) {
	typeOf := reflect.TypeOf(i)
	valueOf := reflect.ValueOf(i)

	v := validator{
		Result: make(Result),
	}

	var err error

	for i := 0; i < typeOf.NumField(); i++ {
		fieldName := typeOf.Field(i).Name
		fieldNameLower := strings.ToLower(fieldName)
		if err := v.validateField(valueOf.Field(i).Interface, fieldName, rules[fieldNameLower]); err != nil {
			return v.Result, err
		}
	}
	return v.Result, err
}

const (
	emptyRulesForField = `no rules found for %s`
)

func (v *validator) validateField(value interface{}, fieldName string, rules []string) error {
	if len(rules) == 0 {
		return fmt.Errorf(emptyRulesForField, fieldName)
	}

	results := []string{}

	var err error

	for _, rule := range rules {
		if rule == "required" {
			if res := r.Required(value); res != "" {
				results = append(results, res)
			}
			continue
		}
		if strings.HasPrefix("max:", rule) {
			maxStrValue := rule[4:]
			maxIntValue, err := strconv.Atoi(maxStrValue)
			if err != nil {
				break
			}
			if res := r.Max(value, maxIntValue); res != "" {
				results = append(results, res)
			}
			continue
		}
		if strings.HasPrefix("min:", rule) {
			minStrValue := rule[4:]
			minIntValue, err := strconv.Atoi(minStrValue)
			if err != nil {
				break
			}
			if res := r.Min(value, minIntValue); res != "" {
				results = append(results, res)
			}
			continue
		}
	}

	if len(results) > 0 {
		v.Result[strings.ToLower(fieldName)] = results
	}

	return err
}
