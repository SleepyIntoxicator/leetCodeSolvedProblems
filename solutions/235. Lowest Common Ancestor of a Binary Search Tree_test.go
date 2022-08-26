package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	fnNames := map[string]func(root, p, q *TreeNode) *TreeNode{
		"lowestCommonAncestor":             lowestCommonAncestor,
		"lowestCommonAncestorFromLeetCode": lowestCommonAncestorFromLeetCode,
	}

	type TestCase struct {
		Name       string
		SourceTree string
		P          int
		Q          int
		Expected   int
	}

	tcs := []TestCase{
		{
			Name:       "1 leetCode",
			SourceTree: "[6,2,8,0,4,7,9,null,null,3,5]",
			P:          2,
			Q:          8,
			Expected:   6,
		},
		{
			Name:       "2 leetCode",
			SourceTree: "[6,2,8,0,4,7,9,null,null,3,5]",
			P:          2,
			Q:          4,
			Expected:   2,
		},
		{
			Name:       "3 leetCode",
			SourceTree: "[2,1]",
			P:          2,
			Q:          1,
			Expected:   2,
		},
		{
			Name:       "4",
			SourceTree: "[6,2,8,0,4,7,9,null,null,3,5]",
			P:          3,
			Q:          5,
			Expected:   4,
		},
		{
			Name:       "5",
			SourceTree: "[6,2,8,0,4,7,9,null,null,3,5]",
			P:          5,
			Q:          0,
			Expected:   2,
		},
		{
			Name:       "6 null",
			SourceTree: "[]",
			P:          1,
			Q:          2,
			Expected:   0,
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				root, err := StringToIntTree(tc.SourceTree)
				assert.NoError(t, err)

				res := fn(root, &TreeNode{Val: tc.P}, &TreeNode{Val: tc.Q})
				if res == nil {
					assert.Equal(t, tc.Expected, 0)
				} else {
					assert.Equal(t, tc.Expected, res.Val)
				}
			})
		}
	}
}
