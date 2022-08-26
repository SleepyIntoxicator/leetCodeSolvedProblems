package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	iterator := root
	for iterator != nil {
		if val == iterator.Val {
			return iterator
		} else if val < iterator.Val {
			iterator = iterator.Left
		} else {
			iterator = iterator.Right
		}
	}

	return nil
}
