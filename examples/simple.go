package examples

import (
	"github.com/fiuskylab/go-valid/validator"
)

type User struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	PhoneNumber          string `json:"phone_number"`
}

func main() {
	u := User{
		Username:             "foo_bar",
		Password:             "secret",
		PasswordConfirmation: "secret",
		PhoneNumber:          "",
	}

	rules := validator.Rules{
		"username": validator.Rule{
			Required: true,
			Min:      4,
			Max:      16,
		},
		"password": validator.Rule{
			Required: true,
			Min:      8,
			Max:      32,
		},
		"passwordconfirmation": validator.Rule{
			Required: true,
			Min:      8,
			Max:      32,
		},
		"phonenumber": validator.Rule{
			Required: false, // Default
			Regex:    `/some_regex/`,
		},
	}

	validator.Check(u, rules)
}
