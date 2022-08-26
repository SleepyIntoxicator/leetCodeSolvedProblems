package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBSTIterative(t *testing.T) {
	fnNames := map[string]func(root *TreeNode) bool{
		"isValidBSTIterative":  isValidBSTIterative,
		"isValidBSTBruteforce": isValidBSTBruteforce,
		"isValidBSTRecursive":  isValidBSTRecursive,
	}

	type TestCase struct {
		Name       string
		SourceTree string
		Expected   bool
	}

	tcs := []TestCase{
		{
			Name:       "1 leetCode",
			SourceTree: "[2,1,3]",
			Expected:   true,
		},
		{
			Name:       "2 leetCode",
			SourceTree: "[5,1,4,null,null,3,6]",
			Expected:   false,
		},
		{
			Name:       "3 leetCode",
			SourceTree: "[5,4,6,null,null,3,7]",
			Expected:   false,
		},
		{
			Name:       "4 empty",
			SourceTree: "[]",
			Expected:   true,
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				root, err := StringToIntTree(tc.SourceTree)
				assert.NoError(t, err)

				res := fn(root)

				assert.Equal(t, tc.Expected, res)
			})
		}
	}
}
