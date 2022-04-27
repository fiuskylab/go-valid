package validator

type (
	// Validator is responsible for
	// handling all validations, and returning
	// the response and/or errors.
	Validator struct {
		isValid          bool
		validationErrors map[string][]string
		rules            Rules
	}

	// Rules stores all rules
	// to validate a struct
	Rules map[string]rule

	rule struct {
		Min *int
		Max *int
	}

	// TODO:
	// 		implement this $H!T
	rulesTypeConstraint interface {
		[]string | []rule
	}
)
