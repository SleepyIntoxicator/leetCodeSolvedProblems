package solutions

import (
	"fmt"
	"testing"

	"github.com/Slintox/leetCodeSolvedProblems/utils"
)

func TestMatrixReshape(t *testing.T) {
	fnNames := map[string]func(mat [][]int, r int, c int) [][]int{
		"matrixReshape":             matrixReshape,
		"matrixReshapeFromLeetCode": matrixReshapeFromLeetCode,
	}

	type TestCase struct {
		Name     string
		Mat      [][]int
		R        int
		C        int
		Expected [][]int
	}

	tcs := []TestCase{
		{
			Name:     "1_leetCode",
			Mat:      [][]int{{1, 2}, {3, 4}},
			R:        1,
			C:        4,
			Expected: [][]int{{1, 2, 3, 4}},
		},
		{
			Name:     "2",
			Mat:      [][]int{{1, 2}, {3, 4}},
			R:        4,
			C:        1,
			Expected: [][]int{{1}, {2}, {3}, {4}},
		},
		{
			Name:     "3_leetCode",
			Mat:      [][]int{{1, 2}, {3, 4}},
			R:        2,
			C:        4,
			Expected: [][]int{{1, 2}, {3, 4}},
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				out := fn(tc.Mat, tc.R, tc.C)

				utils.AssertEqualLists(t, tc.Expected, out)
			})
		}
	}
}
