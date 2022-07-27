package solutions

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	curr := head
	linkHash := make(map[*ListNode]struct{}, 0)
	for curr != nil { // or curr != nil
		if _, ok := linkHash[curr]; ok {
			return true
		} else {
			linkHash[curr] = struct{}{}
		}

		curr = curr.Next
	}

	return false
}
