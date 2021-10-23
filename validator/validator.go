package validator

import (
	"reflect"
	"strings"

	"github.com/fiuskylab/go-valid/validator/rules"
	r "github.com/fiuskylab/go-valid/validator/rules"
)

type validator struct {
	Result Result
}

type Rules map[string]Rule

type Rule struct {
	// Max:
	// 		Lenght of strings, slices, etc.
	// 		Size of number
	Max int
	// Min:
	// 		Lenght of strings, slices, etc.
	// 		Size of number
	Min int
	// Regex is the string of regular expression
	Regex string
	// Type rule
	// compare the received value to the expected type
	Type rules.Type
	// TypeVar rule
	// compare the received type of value from another type of value
	TypeVar interface{}
	// Required
	// 		if the field must have a value
	Required bool
}

func Check(i interface{}, rules Rules) (Result, error) {
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

func (v *validator) validateField(value interface{}, fieldName string, rules Rule) error {
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

	if len(results) > 0 {
		v.Result[strings.ToLower(fieldName)] = results
	}

	return err
}
