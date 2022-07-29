package solutions

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

func LinkedListToSliceInt(head *ListNode) []int {
	if head == nil {
		return nil
	}

	curr := head
	var res []int
	for curr != nil {
		res = append(res, curr.Val)
		curr = curr.Next
	}
	return res
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
