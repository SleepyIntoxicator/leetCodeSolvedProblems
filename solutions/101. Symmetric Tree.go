package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		itemsOnLevel := len(queue)
		for _, v := range queue {
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		j := itemsOnLevel - 1
		for i := 0; i < (itemsOnLevel/2)+1; i++ {
			if !((queue[i].Left != nil && queue[j].Right != nil) || (queue[i].Left == nil && queue[j].Right == nil)) {
				return false
			}
			if queue[i].Val != queue[j].Val {
				return false
			}
			j--
		}

		queue = queue[itemsOnLevel:]
	}

	return true
}

// Used inverted equivalent or XOR formula ¬((¬A∧¬B)∨(A∧B)) or (A∨B)∧¬(A∧B)
// left right  out
//
//	nil  nil  false
//	nil   v   true
//	 v   nil  true
//	 v    v   false
func isSymmetricDFS(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var stackLeft, stackRight []*TreeNode
	currLeft, currRight := root.Left, root.Right

	if !((currLeft != nil && currRight != nil) || (currLeft == nil && currRight == nil)) {
		return false
	}

	for (len(stackLeft) > 0 && len(stackRight) > 0) || (currLeft != nil || currRight != nil) {
		for currLeft != nil && currRight != nil {
			if currLeft.Val != currRight.Val {
				return false
			}
			stackLeft = append(stackLeft, currLeft)
			stackRight = append(stackRight, currRight)
			currLeft = currLeft.Left
			currRight = currRight.Right
		}

		if (currLeft == nil || currRight == nil) && !(currLeft == nil && currRight == nil) {
			return false
		}

		currLeft = stackLeft[len(stackLeft)-1].Right
		currRight = stackRight[len(stackRight)-1].Left
		stackLeft = stackLeft[:len(stackLeft)-1]
		stackRight = stackRight[:len(stackRight)-1]
	}

	return true
}

func isSymmetricRecursive(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return iSRBFS(root.Left, root.Right)
}

func iSRBFS(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val && iSRBFS(left.Left, right.Right) && iSRBFS(left.Right, right.Left)
}
