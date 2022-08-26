package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var pStack, qStack []*TreeNode

	curr := root
	for {
		pStack = append(pStack, curr)
		if p.Val == curr.Val {
			break
		} else if p.Val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	curr = root
	for {
		qStack = append(qStack, curr)
		if q.Val == curr.Val {
			break
		} else if q.Val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	for len(pStack) != len(qStack) {
		if len(pStack) > len(qStack) {
			pStack = pStack[:len(pStack)-1]
		} else {
			qStack = qStack[:len(qStack)-1]
		}
	}

	h := len(pStack) - 1
	for h >= 0 {
		if pStack[h] != qStack[h] {
			h--
		} else {
			break
		}
	}

	return pStack[h]
}

func lowestCommonAncestorFromLeetCode(root, p, q *TreeNode) *TreeNode {
	curr := root
	for curr != nil {
		if p.Val < curr.Val && q.Val < curr.Val {
			curr = curr.Left
		} else if p.Val > curr.Val && q.Val > curr.Val {
			curr = curr.Right
		} else {
			return curr
		}
	}

	return nil
}
