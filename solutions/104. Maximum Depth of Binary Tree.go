package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// using Breadth-First Search
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	max := 0
	level := 0
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		itemsOnLevel := len(queue)
		for i := 0; i < itemsOnLevel; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		level++
		if level > max {
			max = level
		}
		queue = queue[itemsOnLevel:]
	}

	return max
}
