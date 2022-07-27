package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	type TestCase struct {
		Name       string
		RansomNote string
		Magazine   string
		Expected   bool
	}

	tcs := []TestCase{
		{
			Name:       "1_leetCode",
			RansomNote: "a",
			Magazine:   "b",
			Expected:   false,
		},
		{
			Name:       "2_leetCode",
			RansomNote: "aa",
			Magazine:   "ab",
			Expected:   false,
		},
		{
			Name:       "3_leetCode",
			RansomNote: "aa",
			Magazine:   "aab",
			Expected:   true,
		},
		{
			Name:       "4_NotEnoughLengthOfMagazine",
			RansomNote: "aaa",
			Magazine:   "a",
			Expected:   false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, canConstruct(tc.RansomNote, tc.Magazine))
		})
	}
}
