package rules

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRequired(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			require := require.New(t)
			value := 0

			got := Required(value, "value")

			require.NotNil(got)
		})
		t.Run("Not Empty", func(t *testing.T) {
			require := require.New(t)
			value := 1

			got := Required(value, "value")

			require.Nil(got)
		})
	})
}
