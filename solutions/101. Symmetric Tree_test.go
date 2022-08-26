package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSymmetric(t *testing.T) {
	fnNames := map[string]func(root *TreeNode) bool{
		"isSymmetricBFS":       isSymmetric,
		"isSymmetricDFS":       isSymmetricDFS,
		"isSymmetricRecursive": isSymmetricRecursive,
	}

	type TestCase struct {
		Name       string
		SourceTree string
		Expected   bool
	}

	tcs := []TestCase{
		{
			Name:       "Symmetric tree",
			SourceTree: "[1,2,2,3,4,4,3]",
			Expected:   true,
		},
		{
			Name:       "Symmetric long tree",
			SourceTree: "[2,3,3,4,5,5,4,6,null,8,9,9,8,null,6]",
			Expected:   true,
		},
		{
			Name:       "Unsymmetric tree: Inverted child",
			SourceTree: "[1,2,2,null,3,null,3]",
			Expected:   false,
		},
		{
			Name:       "Unsymmetric tree: Different length",
			SourceTree: "[1,2,null]",
			Expected:   false,
		},
		{
			Name:       "Unsymmetric tree: Different length",
			SourceTree: "[1,2,2,null,3]",
			Expected:   false,
		},
		{
			Name:       "Unsymmetric long tree",
			SourceTree: "[2,3,3,4,5,5,4,null,6,8,9,9,8,null,6]",
			Expected:   false,
		},
		{
			Name:       "Unsymmetric tree: Different values",
			SourceTree: "[1,2,2,3,6,9,3]",
			Expected:   false,
		},
		{
			Name:       "Null tree",
			SourceTree: "[]",
			Expected:   false,
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
