package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseListOptimal(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	var revList *ListNode

	for curr != nil {
		revList = &ListNode{
			Val:  curr.Val,
			Next: revList,
		}

		curr = curr.Next
	}

	return revList
}

func reverseListFastSolution(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var tmp []int
	curr := head

	for curr != nil {
		tmp = append(tmp, curr.Val)
		curr = curr.Next
	}

	var revList, revHead *ListNode

	for i := len(tmp) - 1; i >= 0; i-- {
		if revHead == nil {
			revHead = &ListNode{
				Val: tmp[i],
			}
			revList = revHead
		} else {
			revList.Next = &ListNode{
				Val: tmp[i],
			}
			revList = revList.Next
		}
	}

	return revHead
}

func reverseListRecursiveNotOptimal(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &ListNode{
			Val:  head.Val,
			Next: nil,
		}
	}

	revHead := reverseListRecursiveNotOptimal(head.Next)
	revCurr := revHead
	for {
		if revCurr.Next == nil {
			revCurr.Next = &ListNode{
				Val: head.Val,
			}
			break
		}
		revCurr = revCurr.Next
	}

	return revHead
}

func reverseListFromLeetCode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev, curr, next := head, head.Next, head.Next
	head.Next = nil
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
