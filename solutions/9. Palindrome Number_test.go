package solutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fnTestNames := []string{
		"isPalindromeString",
		"isPalindromeBruteforce",
		"isPalindromeOptimal",
	}

	type testCase struct {
		name        string
		in          int
		expectedOut bool
	}
	tcs := []testCase{
		{
			name:        "palindrome with even digits",
			in:          1221,
			expectedOut: true,
		},
		{
			name:        "palindrome with odd digits",
			in:          121,
			expectedOut: true,
		},
		{
			name:        "zero is palindrome",
			in:          0,
			expectedOut: true,
		},
		{
			name:        "not palindrome",
			in:          1222,
			expectedOut: false,
		},
		{
			name:        "negative number is not palindrome",
			in:          -121,
			expectedOut: false,
		},
	}

	for _, fnName := range fnTestNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.name), func(t *testing.T) {
				var out bool

				switch fnName {
				case "isPalindromeString":
					out = IsPalindromeString(tc.in)
				case "isPalindromeBruteforce":
					out = IsPalindromeBruteforce(tc.in)
				case "isPalindromeOptimal":
					out = IsPalindromeOptimal(tc.in)
				}

				assert.Equal(t, tc.expectedOut, out)
			})
		}
	}
}
