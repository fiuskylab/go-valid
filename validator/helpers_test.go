package validator

import "testing"

func TestToLower(t *testing.T) {
	{
		want := "password"
		got := toLower("Password")
		if want != got {
			t.Errorf("Want %s, got %s", want, got)
		}
	}
	{
		want := "password_confirmation"
		got := toLower("PasswordConfirmation")
		if want != got {
			t.Errorf("Want %s, got %s", want, got)
		}
	}
	{
		want := "foo_bar_fuzz"
		got := toLower("FooBarFuzz")
		if want != got {
			t.Errorf("Want %s, got %s", want, got)
		}
	}
}
