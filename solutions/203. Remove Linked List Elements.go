package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	var newList, newListHead *ListNode

	for curr != nil {
		// Adds new elements only if the current value is val
		if curr.Val != val {
			// If the first element
			if newListHead == nil {
				newListHead = &ListNode{
					Val: curr.Val,
				}
				newList = newListHead
				// If the remaining elements
			} else {
				newList.Next = &ListNode{
					Val: curr.Val,
				}
				newList = newList.Next
			}
		}
		curr = curr.Next
	}

	return newListHead
}

func removeElementsFromSource(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	curr := head

	for curr != nil {
		// Delete the next item if next value is val
		if curr.Next != nil && curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	if head.Val == val {
		return head.Next
	}

	return head
}

func removeElementsRecursive(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	head.Next = removeElementsRecursive(head.Next, val)

	if head.Val == val {
		return head.Next
	}

	return head
}
