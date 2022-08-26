package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	fnNames := map[string]func(root *TreeNode) *TreeNode{
		"invertTreeBFS":       invertTreeBFS,
		"invertTreeDFS":       invertTreeDFS,
		"invertTreeRecursive": invertTreeRecursive,
	}

	type TestCase struct {
		Name         string
		SourceTree   string
		ExpectedTree string
	}

	tcs := []TestCase{
		{
			Name:         "Balanced tree",
			SourceTree:   "[4,2,7,1,3,6,9]",
			ExpectedTree: "[4,7,2,9,6,3,1]",
		}, {
			Name:         "Balanced big tree",
			SourceTree:   "[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15]",
			ExpectedTree: "[1,3,2,7,6,5,4,15,14,13,12,11,10,9,8]",
		},
		{
			Name:         "Small tree",
			SourceTree:   "[2,1,3]",
			ExpectedTree: "[2,3,1]",
		},
		{
			Name:         "Small tree with several null child elems",
			SourceTree:   "[2,3,null,1]",
			ExpectedTree: "[2,null,3,null,1]",
		},
		{
			Name:         "Null tree",
			SourceTree:   "[]",
			ExpectedTree: "[]",
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				root, err := StringToIntTree(tc.SourceTree)
				assert.NoError(t, err)

				var res string
				res, err = IntTreeToString(fn(root))
				assert.NoError(t, err)
				assert.Equal(t, tc.ExpectedTree, res)
			})
		}
	}
}
