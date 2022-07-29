package solutions

import (
	"fmt"
	"testing"

	"leetCodeSolvedProblems/utils"
)

func TestReverseList(t *testing.T) {
	fnNames := map[string]func(head *ListNode) *ListNode{
		"reverseListFastSolution":        reverseListFastSolution,
		"reverseListOptimal":             reverseListOptimal,
		"reverseListRecursiveNotOptimal": reverseListRecursiveNotOptimal,
		"reverseListFromLeetCode":        reverseListFromLeetCode,
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
			Src:      []int{1, 2, 3, 4, 5},
			Expected: []int{5, 4, 3, 2, 1},
		},
		{
			Name:     "3_leetCode",
			Src:      []int{1, 2},
			Expected: []int{2, 1},
		},
		{
			Name:     "2_leetCode",
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
