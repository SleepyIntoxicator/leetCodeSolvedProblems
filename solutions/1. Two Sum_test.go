package solutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	fnNames := map[string]func(nums []int, target int) []int{
		"TwoSumNew":        twoSumNew,
		"TwoSumOld":        twoSumOld,
		"TwoSumBruteforce": twoSumBruteforce,
	}

	type TestCase struct {
		Name     string
		Nums     []int
		Target   int
		Expected []int
	}

	tcs := []TestCase{
		{
			Name:     "1",
			Nums:     []int{2, 7, 11, 15},
			Target:   9,
			Expected: []int{0, 1},
		},
		{
			Name:     "2",
			Nums:     []int{3, 2, 4},
			Target:   6,
			Expected: []int{1, 2},
		},
		{
			Name:     "3",
			Nums:     []int{3, 3},
			Target:   6,
			Expected: []int{0, 1},
		},
	}

	for fnName, twoSum := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				out := twoSum(tc.Nums, tc.Target)

				assert.ElementsMatch(t, tc.Expected, out)
			})
		}
	}
}
