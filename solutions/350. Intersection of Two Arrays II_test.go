package solutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersect(t *testing.T) {
	fnNames := map[string]func(nums1 []int, nums2 []int) []int{
		"intersectBruteforce": intersectBruteforce,
	}

	type TestCase struct {
		Name     string
		Nums1    []int
		Nums2    []int
		Expected []int
	}

	tcs := []TestCase{
		{
			Name:     "1",
			Nums1:    []int{1, 2, 2, 1},
			Nums2:    []int{2, 2},
			Expected: []int{2, 2},
		},
		{
			Name:     "2",
			Nums1:    []int{4, 9, 5},
			Nums2:    []int{9, 4, 9, 8, 4},
			Expected: []int{4, 9},
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				out := fn(tc.Nums1, tc.Nums2)

				assert.ElementsMatch(t, tc.Expected, out)
			})
		}
	}
}
