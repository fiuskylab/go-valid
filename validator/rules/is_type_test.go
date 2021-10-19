package rules

import (
	"fmt"
	"testing"
)

func TestIsType(t *testing.T) {
	{
		want := ""
		got := IsType(int64(32), Int64)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
	{
		want := fmt.Sprintf(typeDontMatch, Int.toString())
		got := IsType(uint(12), Int)
		if want != got {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
	}
}
