package solutions

func insertIntoBSTSimple(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	curr := root

	for curr != nil {
		if val < curr.Val {
			if curr.Left != nil {
				curr = curr.Left
			} else {
				curr.Left = &TreeNode{Val: val}
				break
			}
		} else {
			if curr.Right != nil {
				curr = curr.Right
			} else {
				curr.Right = &TreeNode{Val: val}
				break
			}
		}
	}

	return root
}
