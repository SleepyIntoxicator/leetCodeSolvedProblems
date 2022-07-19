package solutions

import (
	"fmt"
	"leetCodeSolvedProblems/utils"
	"testing"
)

func TestMerge(t *testing.T) {
	fnNames := map[string]func(nums1 []int, m int, nums2 []int, n int){
		"merge": merge,
	}

	type TestCase struct {
		Name     string
		Nums1    []int
		m        int
		Nums2    []int
		n        int
		Expected []int
	}

	tcs := []TestCase{
		{
			Name:     "1",
			Nums1:    []int{1, 2, 3, 0, 0, 0},
			m:        3,
			Nums2:    []int{2, 5, 6},
			n:        3,
			Expected: []int{1, 2, 2, 3, 5, 6},
		},
		{
			Name:     "2",
			Nums1:    []int{1},
			m:        1,
			Nums2:    []int{},
			n:        0,
			Expected: []int{1},
		},
		{
			Name:     "3",
			Nums1:    []int{0},
			m:        0,
			Nums2:    []int{1},
			n:        1,
			Expected: []int{1},
		},
		{
			Name:     "4",
			Nums1:    []int{4, 5, 6, 0, 0, 0},
			m:        3,
			Nums2:    []int{1, 2, 3},
			n:        3,
			Expected: []int{1, 2, 3, 4, 5, 6},
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				fn(tc.Nums1, tc.m, tc.Nums2, tc.n)

				utils.AssertEqualLists(t, tc.Expected, tc.Nums1)
			})
		}
	}
}
