package solutions

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	sortedList := &ListNode{}
	head := sortedList

	for {
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				sortedList.Val = list1.Val
				list1 = list1.Next
			} else {
				sortedList.Val = list2.Val
				list2 = list2.Next
			}
		} else if list1 != nil {
			sortedList.Val = list1.Val
			list1 = list1.Next
		} else if list2 != nil {
			sortedList.Val = list2.Val
			list2 = list2.Next
		}

		if list1 == nil && list2 == nil {
			break
		}

		sortedList.Next = &ListNode{}
		sortedList = sortedList.Next
	}

	return head
}

func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	sorted := &ListNode{}

	if list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			sorted.Val = list1.Val
			sorted.Next = mergeTwoListsRecursive(list1.Next, list2)
		} else {
			sorted.Val = list2.Val
			sorted.Next = mergeTwoListsRecursive(list1, list2.Next)
		}
	} else if list1 != nil {
		sorted.Val = list1.Val
		sorted.Next = mergeTwoListsRecursive(list1.Next, nil)
	} else if list2 != nil {
		sorted.Val = list2.Val
		sorted.Next = mergeTwoListsRecursive(nil, list2.Next)
	}

	return sorted
}
