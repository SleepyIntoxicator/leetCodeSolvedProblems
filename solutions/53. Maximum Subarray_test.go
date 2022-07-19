package solutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	fnNames := map[string]func([]int) int{
		"maxSubArrayShort": maxSubArrayShort,
		"maxSubArrayFirst": maxSubArrayFirst,
	}

	type TestCase struct {
		Name     string
		Nums     []int
		Expected int
	}

	tcs := []TestCase{
		{
			Name:     "long sequence",
			Nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			Expected: 6,
		},
		{
			Name:     "accounting for the last numbers",
			Nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4, 6},
			Expected: 11,
		},
		{
			Name:     "one number",
			Nums:     []int{1},
			Expected: 1,
		},
		{
			Name:     "subarr len is equal to arr len",
			Nums:     []int{5, 4, -1, 7, 8},
			Expected: 23,
		},
		{
			Name:     "left or right",
			Nums:     []int{5, 4, -8, 7, 8},
			Expected: 16,
		},
		{
			Name:     "left or right 2",
			Nums:     []int{5, 4, -9, 7, 8},
			Expected: 15,
		},
		{
			Name:     "all positive",
			Nums:     []int{1, 2, 3},
			Expected: 6,
		},
		{
			Name:     "all negative",
			Nums:     []int{-1, -2, -3},
			Expected: -1,
		},
	}

	for fnName, maxSubArray := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				out := maxSubArray(tc.Nums)

				assert.Equal(t, tc.Expected, out)
			})

		}
	}
}
