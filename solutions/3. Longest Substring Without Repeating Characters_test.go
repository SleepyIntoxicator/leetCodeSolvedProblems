package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type TestCase struct {
		Name        string
		InpString   string
		ExpectedLen int
	}

	tcs := []TestCase{
		{
			Name:        "1 leetCode",
			InpString:   "abcabcbb",
			ExpectedLen: 3,
		},
		{
			Name:        "2 leetCode",
			InpString:   "bbbbb",
			ExpectedLen: 1,
		},
		{
			Name:        "3 leetCode",
			InpString:   "pwwkew",
			ExpectedLen: 3,
		},
		{
			Name:        "4 leetCode",
			InpString:   " ",
			ExpectedLen: 1,
		},
		{
			Name:        "5 leetCode",
			InpString:   "p",
			ExpectedLen: 1,
		},
		{
			Name:        "6 leetCode",
			InpString:   " p",
			ExpectedLen: 2,
		},
		{
			Name:        "7 leetCode",
			InpString:   "dvdf",
			ExpectedLen: 3,
		},
		{
			Name:        "8 leetCode",
			InpString:   "aabaab!bb",
			ExpectedLen: 3,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.ExpectedLen, lengthOfLongestSubstring(tc.InpString))
		})
	}
}
