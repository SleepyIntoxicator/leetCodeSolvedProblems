package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertIntoBSTSimple(t *testing.T) {
	fnNames := map[string]func(root *TreeNode, val int) *TreeNode{
		"insertIntoBSTSimple": insertIntoBSTSimple,
	}

	type TestCase struct {
		Name          string
		SourceTree    string
		Val           int
		ExpectedTrees []string
	}

	tcs := []TestCase{
		{
			Name:       "1",
			SourceTree: "[4,2,7,1,3]",
			Val:        5,
			ExpectedTrees: []string{
				"[4,2,7,1,3,5]",
				"[5,2,7,1,3,null,null,null,null,null,4]",
			},
		},
		{
			Name:       "2",
			SourceTree: "[40,20,60,10,30,50,70]",
			Val:        25,
			ExpectedTrees: []string{
				"[40,20,60,10,30,50,70,null,null,25]",
			},
		},
		{
			Name:       "3",
			SourceTree: "[4,2,7,1,3]",
			Val:        5,
			ExpectedTrees: []string{
				"[4,2,7,1,3,5]",
			},
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				root, err := StringToIntTree(tc.SourceTree)
				assert.NoError(t, err)

				var out string
				out, err = IntTreeToString(fn(root, tc.Val))
				assert.NoError(t, err)
				assert.Contains(t, tc.ExpectedTrees, out)
			})
		}
	}
}
