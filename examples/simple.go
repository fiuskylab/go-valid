package examples

import (
	"github.com/fiuskylab/go-valid/validator"
	"github.com/fiuskylab/go-valid/validator/rules"
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

	rules := rules.Rules{
		"username": rules.Rule{
			Required: true,
			Min:      4,
			Max:      16,
		},
		"password": rules.Rule{
			Required: true,
			Min:      8,
			Max:      32,
		},
		"passwordconfirmation": rules.Rule{
			Required: true,
			Min:      8,
			Max:      32,
		},
		"phonenumber": rules.Rule{
			Required: false, // Default
			Regex:    `/some_regex/`,
		},
	}

	validator.Check(u, rules)
}
