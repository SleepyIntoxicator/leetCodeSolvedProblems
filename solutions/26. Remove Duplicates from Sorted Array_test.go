package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	fnNames := map[string]func(nums []int) int{
		"RemoveDuplicates":             RemoveDuplicates,
		"RemoveDuplicatesFromLeetCode": RemoveDuplicatesFromLeetCode,
	}

	type TestCase struct {
		Name     string
		Array    []int
		Expected []int
	}

	tcs := []TestCase{
		{
			Name:     "LeetCode_1",
			Array:    []int{1, 1, 2},
			Expected: []int{1, 2},
		},
		{
			Name:     "LeetCode_2",
			Array:    []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			Expected: []int{0, 1, 2, 3, 4},
		},
		{
			Name:     "No_duplicates",
			Array:    []int{1, 2, 3, 4, 5, 6, 7, 8},
			Expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			Name:     "Empty array",
			Array:    []int{},
			Expected: []int{},
		},
		{
			Name:     "One value",
			Array:    []int{1},
			Expected: []int{1},
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				arr := make([]int, len(tc.Array))
				copy(arr, tc.Array)
				out := fn(arr)

				assert.Equal(t, tc.Expected, arr[:out])
			})
		}
	}
}
