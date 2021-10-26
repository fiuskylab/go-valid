package rules

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
	Type Type
	// TypeVar rule
	// compare the received type of value from another type of value
	TypeVar interface{}
	// Required
	// 		if the field must have a value
	Required bool

	// In
	// 		check if a value is inside of a slice of values
	In []interface{}

	// If must have a confirmation field
	// e.g user and userconfirmation
	Confirmed bool
}
