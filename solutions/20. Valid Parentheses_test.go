package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	type TestCase struct {
		Name     string
		Input    string
		Expected bool
	}

	tcs := []TestCase{
		{
			Name:     "1_leetCode",
			Input:    "()",
			Expected: true,
		},
		{
			Name:     "2_leetCode",
			Input:    "()[]{}",
			Expected: true,
		},
		{
			Name:     "3_leetCode",
			Input:    "(]",
			Expected: false,
		},
		{
			Name:     "4",
			Input:    "(())",
			Expected: true,
		},
		{
			Name:     "5",
			Input:    "))",
			Expected: false,
		},
		{
			Name:     "6",
			Input:    "((",
			Expected: false,
		},
		{
			Name:     "7",
			Input:    "[",
			Expected: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			out := isValid(tc.Input)

			assert.Equal(t, tc.Expected, out)
		})
	}
}
