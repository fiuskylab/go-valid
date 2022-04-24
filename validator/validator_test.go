package validator

import (
	"reflect"
	"testing"

	"github.com/fiuskylab/go-valid/validator/rules"
)

type TestStruct struct {
	String   string
	Int      int
	SliceStr []string
}

func TestCheck(t *testing.T) {
	ts := TestStruct{
		String:   "foo",
		Int:      10,
		SliceStr: []string{"foo", "bar"},
	}

	rules := rules.Rules{
		"string": rules.Rule{
			Required: true,
			Min:      2,
			Max:      10,
		},
		"int": rules.Rule{
			Required: true,
			Min:      2,
			Max:      13,
		},
		"slicestr": rules.Rule{
			Required: true,
			Min:      2,
			Max:      10,
		},
	}

	{
		want := Result{}
		got, gotErr := Check(ts, rules)

		if !reflect.DeepEqual(want, got) {
			t.Errorf(`want "%s", got "%s"`, want, got)
		}
		if gotErr != nil {
			t.Errorf(`Error must be nil, received %s`, gotErr.Error())
		}
	}
}
