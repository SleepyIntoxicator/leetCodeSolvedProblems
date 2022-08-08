package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		itemsOnLevel := len(queue)
		if itemsOnLevel != 1 && (itemsOnLevel%2 == 1) {
			return false
		}
		j := itemsOnLevel - 1
		for i := 0; i < (itemsOnLevel / 2); i++ {

			if !((queue[i].Left != nil && queue[j].Right != nil) || (queue[i].Left == nil && queue[j].Right == nil)) {
				return false
			}
			if queue[i].Val != queue[j].Val {
				return false
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			j--
		}

		queue = queue[itemsOnLevel:]
	}

	return true
}

func isSymmetricD(root *TreeNode) bool {
	if root == nil {
		return false
	}

	var sl, sr []*TreeNode
	cl, cr := root.Left, root.Right

	if !((cl != nil && cr != nil) || (cl == nil && cr == nil)) {
		return false
	}

	for (len(sl) > 0 && len(sr) > 0) && (cl != nil && cr != nil) {
		for cl != nil && cr != nil {
			if cl.Val != cr.Val {
				return false
			}
			sl = append(sl, cl)
			sr = append(sr, cr)
			cl = cl.Left
			cr = cr.Right
		}
		if (cl == nil || cr == nil) && !(cl == nil && cr == nil) {
			return false
		}
		if len(sl) != len(sr) {
			return false
		}

		cl = sl[len(sl)-1].Right
		cr = sr[len(sr)-1].Left
		sl = sl[:len(sl)-1]
		sr = sr[:len(sr)-1]
		if cl.Val != cr.Val {
			return false
		}
	}

	if len(sl) > 0 || len(sr) > 0 {
		return false
	}

	return true
}
