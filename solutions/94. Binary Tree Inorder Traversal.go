package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var r []int
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		r = append(r, curr.Val)
		curr = curr.Right
	}

	return r
}

func recursiveTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	if root.Left != nil {
		res = recursiveTraversal(root.Left)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		res = append(res, recursiveTraversal(root.Right)...)
	}
	return res
}
