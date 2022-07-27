package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	type TestCase struct {
		Name     string
		Src      []int
		Head     *ListNode
		Pos      int
		Expected bool
	}

	tcs := []TestCase{
		{
			Name:     "1_leetCode",
			Src:      []int{3, 2, 0, 4},
			Pos:      1,
			Expected: true,
		}, {
			Name:     "2_leetCode",
			Src:      []int{1, 2},
			Pos:      0,
			Expected: true,
		}, {
			Name:     "3_leetCode",
			Src:      []int{1},
			Pos:      -1,
			Expected: false,
		},
	}

	for i := range tcs {
		tcs[i].Head = SliceToLinkedListInt(tcs[i].Src)
		PrepareLinkedListCycle(tcs[i].Head, tcs[i].Pos)
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expected, hasCycle(tc.Head))
		})
	}
}

// TODO: use generics
func SliceToLinkedListInt(src []int) *ListNode {
	if len(src) == 0 {
		return nil
	}

	curr := &ListNode{}
	head := curr

	for i, v := range src {
		curr.Val = v
		if i < len(src)-1 {
			curr.Next = &ListNode{}
		}
		curr = curr.Next
	}

	return head
}

func PrepareLinkedListCycle(head *ListNode, pos int) {
	if head == nil || pos == -1 {
		return
	}

	nodeCounter := 0
	var posNode *ListNode

	curr := head
	for curr != nil {
		if nodeCounter == pos {
			posNode = curr
		}
		if curr.Next == nil {
			curr.Next = posNode
			break
		}
		nodeCounter++
		curr = curr.Next
	}
}
