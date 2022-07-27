package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	type TestCase struct {
		Name     string
		Nums     []int
		Target   int
		Expected int
	}

	tcs := []TestCase{
		{
			Name:     "search in even number of numbers",
			Nums:     []int{1, 2, 3, 4, 5, 6, 7, 8},
			Target:   5,
			Expected: 4,
		},
		{
			Name:     "search in odd number of numbers",
			Nums:     []int{1, 2, 3, 4, 5, 6, 7},
			Target:   4,
			Expected: 3,
		},
		{
			Name:     "search for a missing number",
			Nums:     []int{1, 2, 3, 4, 6, 7, 8},
			Target:   5,
			Expected: -1,
		},
		{
			Name:     "search in a null array of numbers",
			Nums:     []int{},
			Target:   0,
			Expected: -1,
		},
	}

	for _, tc := range tcs {
		t.Run("704. BinarySearch/"+tc.Name, func(t *testing.T) {
			out := Search(tc.Nums, tc.Target)

			assert.Equal(t, tc.Expected, out)
		})
	}
}
