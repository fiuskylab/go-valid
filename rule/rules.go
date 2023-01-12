package rule

// Rules - datastructure to store every rule to
// validate a given value.
// If a rule is nil(not set), it will be ignored.
type Rules struct {
	// Required - Checks if the value is empty/nil.
	Required *bool
	// Min -
	// 	- For numbers it checks if it's greater than given value.
	// 	- For slices it checks if it's length is greater than given value.
	Max *int
	// Min -
	// 	- For numbers it checks if it's less than given value.
	// 	- For slices it checks if it's length is less than given value.
	Min *int
}
