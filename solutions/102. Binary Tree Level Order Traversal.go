package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var r [][]int

	queue := []*TreeNode{root}
	var level int

	for len(queue) > 0 {
		itemsOnLevel := len(queue)
		r = append(r, []int{})
		for i := 0; i < itemsOnLevel; i++ {
			r[level] = append(r[level], queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		level++
		queue = queue[itemsOnLevel:]
	}

	return r
}
