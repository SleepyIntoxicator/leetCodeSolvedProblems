package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	type testCase struct {
		name        string
		in          []string
		expectedOut string
	}

	tcs := []testCase{
		{
			name: "normal prefix: fl",
			in: []string{
				"flower",
				"flow",
				"flight",
			},
			expectedOut: "fl",
		},
		{
			name: "last longest common prefix",
			in: []string{
				"synchromoment",
				"synchronization",
				"synchrophasotron",
			},
			expectedOut: "synchro",
		},
		{
			name: "without common prefix",
			in: []string{
				"dog",
				"racecar",
				"car",
			},
			expectedOut: "",
		},
		{
			name: "with empty line",
			in: []string{
				"dog",
				"",
				"racecar",
				"car",
			},
			expectedOut: "",
		},
		{
			name: "with empty line",
			in: []string{
				"lizzarddogo",
				"lizzarddogo",
				"lizzarddogo",
			},
			expectedOut: "lizzarddogo",
		},
	}

	for _, tc := range tcs {
		t.Run("14. LongestCommonPrefix/"+tc.name, func(t *testing.T) {
			out := LongestCommonPrefix(tc.in)

			assert.Equal(t, tc.expectedOut, out)
		})
	}
}
