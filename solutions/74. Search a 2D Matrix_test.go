package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrix(t *testing.T) {
	type TestCase struct {
		Name     string
		Matrix   [][]int
		Target   int
		Expected bool
	}

	tcs := []TestCase{
		{
			Name:     "1_leetCode",
			Matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			Target:   3,
			Expected: true,
		},
		{
			Name:     "2_leetCode",
			Matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			Target:   13,
			Expected: false,
		},
		{
			Name:     "3_mustFoundInRightPart",
			Matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 62, 63, 64}, {66, 67, 79, 82}},
			Target:   79,
			Expected: true,
		},
		{
			Name:     "4_MustFoundInZeroColumn",
			Matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}, {61, 62, 63, 64}, {66, 67, 79, 82}},
			Target:   10,
			Expected: true,
		},
		{
			Name:     "5_NotFound",
			Matrix:   [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}},
			Target:   100,
			Expected: false,
		},
		{
			Name:     "6_matrix_1x1",
			Matrix:   [][]int{{1}},
			Target:   100,
			Expected: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, searchMatrix(tc.Matrix, tc.Target))
		})
	}
}
