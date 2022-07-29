package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	var newHead, newList *ListNode
	prev := -10000

	for curr != nil {
		if curr.Val != prev {
			if newHead == nil {
				newHead = &ListNode{
					Val: curr.Val,
				}
				newList = newHead
			} else {
				newList.Next = &ListNode{
					Val: curr.Val,
				}
				newList = newList.Next
			}
		}
		prev = curr.Val
		curr = curr.Next
	}

	return newHead
}
