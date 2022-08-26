package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBSTIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var stack []*TreeNode
	max := -(1 << 32)
	curr := root

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		if curr.Val > max {
			max = curr.Val
		} else {
			return false
		}
		stack = stack[:len(stack)-1]
		curr = curr.Right
	}

	return true
}

func isValidBSTBruteforce(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var stack []*TreeNode
	var intStack []int
	curr := root

	for len(stack) > 0 || curr != nil {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		intStack = append(intStack, curr.Val)
		stack = stack[:len(stack)-1]
		curr = curr.Right
	}

	for i := 1; i < len(intStack); i++ {
		if intStack[i] <= intStack[i-1] {
			return false
		}
	}

	return true
}

func isValidBSTRecursive(root *TreeNode) bool {
	if root == nil {
		return true
	}
	max := new(int)
	*max = -(1 << 32)

	return iVBSTR(root, max)
}

func iVBSTR(root *TreeNode, max *int) bool {
	var r bool

	if root.Left != nil {
		r = iVBSTR(root.Left, max)
		if r == false {
			return false
		}
	}

	if root.Val <= *max {
		return false
	}
	*max = root.Val

	if root.Right != nil {
		r = iVBSTR(root.Right, max)
		if r == false {
			return false
		}
	}

	return true
}
