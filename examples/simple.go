package examples

import "github.com/fiuskylab/go-valid/validator"

type User struct {
	Username             string `v:"username"`
	Password             string `v:"password"`
	PasswordConfirmation string `v:"password_confirmation"`
	Email                string `v:"email"`
}

const (
	userRules = validator.Rules{
		"username": {
			Required: true,
			Min:      6,
			Max:      26,
		},
		"password": {
			Required:  true,
			Min:       8,
			Max:       32,
			Confirmed: true,
		},
		"email": {
			Required: true,
			Regex:    "some_regex",
		},
	}
)

func main() {
	_ = validator.Rules{
		"key": {
			Min: validator.ToPtr[uint](1),
		},
	}

	u := User{
		Username:             "rafiusky",
		Password:             "something",
		PasswordConfirmation: "something",
		Email:                "rafiusky@email.com",
	}

	// Here the error is for internal errors,
	// not validation ones.
	res, err := userRules.Validate(u)

	// Returns a boolean, if the
	// struct is valid or not
	if !res.Valid() {
		// handles if not valid

		// MapResults returns the validation errors
		// in the format of a map, so it can be returned
		// in a JSON format.
		m := res.MapResults()
	}
}
