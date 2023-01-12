package rules

import (
	"fmt"
)

const (
	requiredError = "'%s' is required"
)

// Required checks if given value is zero or not.
func Required[T comparable](v T, field string) error {
	err := fmt.Errorf(requiredError, field)

	if v == *(new(T)) {
		return err
	}

	return nil
}
