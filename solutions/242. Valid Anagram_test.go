package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	type TestCase struct {
		Name     string
		S        string
		T        string
		Expected bool
	}

	tcs := []TestCase{
		{
			Name:     "1",
			S:        "anagram",
			T:        "nagaram",
			Expected: true,
		},
		{
			Name:     "2",
			S:        "rat",
			T:        "car",
			Expected: false,
		},
		{
			Name:     "3",
			S:        "ratata",
			T:        "caro",
			Expected: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, isAnagram(tc.S, tc.T))
		})
	}
}
