package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPathSum(t *testing.T) {
	fnNames := map[string]func(root *TreeNode, targetSum int) bool{
		"hasPathSum":                   hasPathSum,
		"hasPathSumRecursive":          hasPathSumRecursive,
		"hasPathSumRecursiveOptimized": hasPathSumRecursiveOptimized,
	}

	type TestCase struct {
		Name       string
		SourceTree string
		TargetSum  int
		Expected   bool
	}

	tcs := []TestCase{
		{
			Name:       "Has path sum: left-right",
			SourceTree: "[5,4,8,11,null,13,4,7,2,null,null,null,1]",
			TargetSum:  22,
			Expected:   true,
		},
		{
			Name:       "Has path sum: right-left",
			SourceTree: "[5,4,8,11,null,13,4,7,2,null,null,null,1]",
			TargetSum:  26,
			Expected:   true,
		},
		{
			Name:       "Has path sum: right-right",
			SourceTree: "[5,4,8,11,null,13,4,7,2,null,null,null,1]",
			TargetSum:  18,
			Expected:   true,
		},
		{
			Name:       "Has not path sum",
			SourceTree: "[1,2,3]",
			TargetSum:  5,
			Expected:   false,
		},
		{
			Name:       "Empty tree",
			SourceTree: "[]",
			TargetSum:  0,
			Expected:   false,
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				root, err := StringToIntTree(tc.SourceTree)
				assert.NoError(t, err)

				res := fn(root, tc.TargetSum)

				assert.Equal(t, tc.Expected, res)
			})
		}
	}
}
