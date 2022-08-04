package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var stack []*TreeNode
	var r []int

	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			r = append(r, curr.Val)
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr = curr.Right
	}

	return r
}

func preorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var r []int

	if root.Left != nil {
		r = preorderTraversalRecursive(root.Left)
	}
	r = append(r, root.Val)
	if root.Right != nil {
		r = append(r, preorderTraversalRecursive(root.Right)...)
	}

	return r
}
