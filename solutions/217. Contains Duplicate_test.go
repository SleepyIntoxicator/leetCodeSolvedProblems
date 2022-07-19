package solutions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsDuplicate(t *testing.T) {
	fnNames := map[string]func([]int) bool{
		"ContainsDuplicateWithHashMap": containsDuplicateWithHashMap,
		"ContainsDuplicateWithSort":    containsDuplicateWithSort,
	}

	type TestCase struct {
		Name     string
		Nums     []int
		Expected bool
	}

	tcs := []TestCase{
		{
			Name:     "Contains duplicate",
			Nums:     []int{1, 2, 3, 1},
			Expected: true,
		},
		{
			Name:     "No contains duplicate",
			Nums:     []int{1, 2, 3, 4},
			Expected: false,
		},
		{
			Name:     "Lots of duplicates",
			Nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			Expected: true,
		},
	}

	for fnName, containsDuplicate := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("ContainsDuplicate/%s/%s", fnName, tc.Name), func(t *testing.T) {

				out := containsDuplicate(tc.Nums)

				assert.Equal(t, tc.Expected, out)
			})
		}
	}
}
