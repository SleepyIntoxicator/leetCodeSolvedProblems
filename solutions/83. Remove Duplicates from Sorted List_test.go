package solutions

import (
	"fmt"
	"testing"

	"leetCodeSolvedProblems/utils"
)

func TestDeleteDuplicates(t *testing.T) {
	fnNames := map[string]func(head *ListNode) *ListNode{
		"deleteDuplicates": deleteDuplicates,
	}

	type TestCase struct {
		Name     string
		Src      []int
		Head     *ListNode
		Expected []int
	}

	tcs := []TestCase{
		{
			Name:     "1_leetCode",
			Src:      []int{1, 1, 2},
			Expected: []int{1, 2},
		},
		{
			Name:     "2_leetCode",
			Src:      []int{1, 1, 2, 3, 3},
			Expected: []int{1, 2, 3},
		},
		{
			Name:     "3",
			Src:      []int{1, 2, 3, 4, 5, 6},
			Expected: []int{1, 2, 3, 4, 5, 6},
		},
		{
			Name:     "4_empty",
			Src:      []int{},
			Expected: []int{},
		},
	}

	for i := range tcs {
		tcs[i].Head = SliceToLinkedListInt(tcs[i].Src)
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				out := LinkedListToSliceInt(fn(tc.Head))

				utils.AssertEqualLists(t, tc.Expected, out)
			})
		}
	}
}
