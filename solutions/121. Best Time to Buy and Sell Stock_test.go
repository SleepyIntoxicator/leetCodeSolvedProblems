package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
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
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			out := maxProfit(tc.Prices)

			assert.Equal(t, tc.Expected, out)
		})
	}
}
