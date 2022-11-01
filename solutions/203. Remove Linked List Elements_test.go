package solutions

import (
	"fmt"
	"testing"

	"github.com/Slintox/leetCodeSolvedProblems/utils"
)

func TestRemoveElements(t *testing.T) {
	fnNames := map[string]func(head *ListNode, val int) *ListNode{
		"removeElements":           removeElements,
		"removeElementsFromSource": removeElementsFromSource,
		"removeElementsRecursive":  removeElementsRecursive,
	}

	type TestCase struct {
		Name     string
		Src      []int
		Head     *ListNode
		Val      int
		Expected []int
	}

	tcs := []TestCase{
		{
			Name:     "1_leetCode",
			Src:      []int{1, 2, 6, 3, 4, 5, 6},
			Val:      6,
			Expected: []int{1, 2, 3, 4, 5},
		},
		{
			Name:     "2_leetCode",
			Src:      []int{},
			Val:      1,
			Expected: []int{},
		},
		{
			Name:     "3_leetCode",
			Src:      []int{7, 7, 7, 7},
			Val:      7,
			Expected: []int{},
		},
	}

	for i := range tcs {
		tcs[i].Head = SliceToLinkedListInt(tcs[i].Src)
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				out := LinkedListToSliceInt(fn(tc.Head, tc.Val))

				utils.AssertEqualLists(t, tc.Expected, out)
			})
		}
	}
}
