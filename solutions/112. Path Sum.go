package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// Breaks the original tree
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	var stack []*TreeNode
	curr := root
	currSum := 0

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			currSum += curr.Val
			if currSum == targetSum && curr.Left == nil && curr.Right == nil {
				return true
			}

			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		currSum -= curr.Val
		if curr.Right != nil {
			curr.Right.Val += curr.Val
		}
		stack = stack[:len(stack)-1]
		curr = curr.Right
	}

	return false
}

func hasPathSumRecursive(root *TreeNode, targetSum int) bool {
	return hPSR(root, 0, targetSum)
}

func hPSR(root *TreeNode, currSum, targetSum int) bool {
	if root == nil {
		return false
	}

	currSum += root.Val
	if root.Left == nil && root.Right == nil {
		return currSum == targetSum
	}
	if hPSR(root.Left, currSum, targetSum) || hPSR(root.Right, currSum, targetSum) {
		return true
	}
	return false
}

// With optimization from leetCode
func hasPathSumRecursiveOptimized(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	if hasPathSumRecursiveOptimized(root.Left, targetSum-root.Val) || hasPathSumRecursiveOptimized(root.Right, targetSum-root.Val) {
		return true
	}
	return false
}
