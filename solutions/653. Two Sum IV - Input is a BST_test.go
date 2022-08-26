package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindTarget(t *testing.T) {
	fnNames := map[string]func(root *TreeNode, targetSum int) bool{
		"findTarget": findTarget,
	}

	type TestCase struct {
		Name       string
		SourceTree string
		TargetSum  int
		Expected   bool
	}

	tcs := []TestCase{
		{
			Name:       "1 leet code",
			SourceTree: "[5,3,6,2,4,null,7]",
			TargetSum:  9,
			Expected:   true,
		},
		{
			Name:       "2 leet code",
			SourceTree: "[5,3,6,2,4,null,7]",
			TargetSum:  28,
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
