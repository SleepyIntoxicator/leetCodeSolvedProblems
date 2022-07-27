package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	fnNames := map[string]func(prices []int) int{
		"maxProfitBruteforce": maxProfitBruteforce,
		"maxProfit":           maxProfit,
	}

	type TestCase struct {
		Name     string
		Prices   []int
		Expected int
	}

	tcs := []TestCase{
		{
			Name:     "1",
			Prices:   []int{7, 1, 5, 3, 6, 4},
			Expected: 5,
		},
		{
			Name:     "2",
			Prices:   []int{7, 6, 4, 3, 1},
			Expected: 0,
		},
		{
			Name:     "3",
			Prices:   []int{1, 5, 9, 6, 3, 0},
			Expected: 8,
		},
		{
			Name:     "4",
			Prices:   []int{2, 1, 2, 0, 1},
			Expected: 1,
		},
		{
			Name:     "4_leetCode",
			Prices:   []int{2, 1, 2, 1, 0, 1, 2},
			Expected: 2,
		},
		{
			Name:     "4_special",
			Prices:   []int{7, 1, 5, 3, 0, 4},
			Expected: 4,
		},
		{
			Name:     "5",
			Prices:   []int{5, 3, 2, 0, 1},
			Expected: 1,
		},
		{
			Name:     "6_leetCode",
			Prices:   []int{3, 3, 5, 0, 0, 3, 1, 4},
			Expected: 4,
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				out := fn(tc.Prices)

				assert.Equal(t, tc.Expected, out)
			})
		}
	}
}
