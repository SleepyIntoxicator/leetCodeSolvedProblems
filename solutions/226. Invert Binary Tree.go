package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTreeBFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		itemsOnLevel := len(queue)
		for i := 0; i < itemsOnLevel; i++ {
			queue[i].Left, queue[i].Right = queue[i].Right, queue[i].Left

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		queue = queue[itemsOnLevel:]
	}

	return root
}

func invertTreeDFS(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var stack []*TreeNode
	curr := root

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			curr.Left, curr.Right = curr.Right, curr.Left
			stack = append(stack, curr)
			curr = curr.Right
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr = curr.Left
	}

	return root
}

func invertTreeRecursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = iTR(root.Left, root.Right)

	return root
}

func iTR(left, right *TreeNode) (*TreeNode, *TreeNode) {
	if left == nil && right == nil {
		return nil, nil
	}

	if left != nil {
		left.Left, left.Right = iTR(left.Left, left.Right)
	}
	if right != nil {
		right.Left, right.Right = iTR(right.Left, right.Right)
	}

	return right, left
}
