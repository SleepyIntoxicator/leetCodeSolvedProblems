package solutions

import (
	"testing"

	"github.com/Slintox/leetCodeSolvedProblems/utils"
)

func TestPascalsTriangle(t *testing.T) {
	type TestCase struct {
		Name     string
		NumRows  int
		Expected [][]int
	}

	tcs := []TestCase{
		{
			Name:     "1",
			NumRows:  1,
			Expected: [][]int{{1}},
		},
		{
			Name:     "2",
			NumRows:  2,
			Expected: [][]int{{1}, {1, 1}},
		},
		{
			Name:     "3",
			NumRows:  3,
			Expected: [][]int{{1}, {1, 1}, {1, 2, 1}},
		},
		{
			Name:     "4",
			NumRows:  4,
			Expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}},
		},
		{
			Name:     "5",
			NumRows:  5,
			Expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			out := generate(tc.NumRows)

			utils.AssertEqualLists(t, tc.Expected, out)
		})
	}
}
