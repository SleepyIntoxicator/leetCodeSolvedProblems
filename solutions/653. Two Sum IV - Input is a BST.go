package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}

	hashSum := make(map[int]struct{}, 0)
	var stack []*TreeNode
	curr := root

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		if _, ok := hashSum[curr.Val]; ok {
			return true
		} else {
			hashSum[k-curr.Val] = struct{}{}
		}
		stack = stack[:len(stack)-1]
		curr = curr.Right
	}

	return false
}
