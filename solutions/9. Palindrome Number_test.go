package solutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	fnTestNames := map[string]func(int) bool{
		"isPalindromeString":     isPalindromeString,
		"isPalindromeBruteforce": isPalindromeBruteforce,
		"isPalindromeOptimal":    isPalindromeOptimal,
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

	for fnName, isPalindrome := range fnTestNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("9. PalindromeNumber/%s/%s", fnName, tc.name), func(t *testing.T) {

				out := isPalindrome(tc.in)

				assert.Equal(t, tc.expectedOut, out)
			})
		}
	}
}
