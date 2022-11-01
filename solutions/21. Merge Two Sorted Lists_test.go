package solutions

import (
	"fmt"
	"testing"

	"github.com/Slintox/leetCodeSolvedProblems/utils"
)

func TestMergeTwoLists(t *testing.T) {
	fnNames := map[string]func(list1 *ListNode, list2 *ListNode) *ListNode{
		"mergeTwoLists":          mergeTwoLists,
		"mergeTwoListsRecursive": mergeTwoListsRecursive,
	}

	type TestCase struct {
		Name     string
		SrcList1 []int
		SrcList2 []int
		List1    *ListNode
		List2    *ListNode
		Expected []int
	}

	tcs := []TestCase{
		{
			Name:     "1_leetCode",
			SrcList1: []int{1, 2, 4},
			SrcList2: []int{1, 3, 4},
			Expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			Name:     "2_leetCode",
			SrcList1: []int{},
			SrcList2: []int{},
			Expected: []int{},
		},
		{
			Name:     "3_leetCode",
			SrcList1: []int{},
			SrcList2: []int{0},
			Expected: []int{0},
		},
	}

	for i := range tcs {
		tcs[i].List1 = SliceToLinkedListInt(tcs[i].SrcList1)
		tcs[i].List2 = SliceToLinkedListInt(tcs[i].SrcList2)
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				out := LinkedListToSliceInt(fn(tc.List1, tc.List2))

				utils.AssertEqualLists(t, tc.Expected, out)
			})
		}
	}
}
