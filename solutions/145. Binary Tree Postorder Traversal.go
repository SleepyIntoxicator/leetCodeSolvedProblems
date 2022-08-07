package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var r []int
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			r = append([]int{curr.Val}, r...)
			stack = append(stack, curr)
			curr = curr.Right
		}

		curr = stack[len(stack)-1].Left
		stack = stack[:len(stack)-1]
	}

	return r
}

func postorderTraversalRecursive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var res []int
	if root.Left != nil {
		res = postorderTraversalRecursive(root.Left)
	}
	if root.Right != nil {
		res = append(res, postorderTraversalRecursive(root.Right)...)
	}
	res = append(res, root.Val)
	return res
}
