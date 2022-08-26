package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchBST(t *testing.T) {
	type TestCase struct {
		Name         string
		SourceTree   string
		SearchVal    int
		ExpectedTree string
	}

	tcs := []TestCase{
		{
			Name:         "Search: left",
			SourceTree:   "[4,2,7,1,3]",
			SearchVal:    2,
			ExpectedTree: "[2,1,3]",
		},
		{
			Name:         "Search: not existed",
			SourceTree:   "[4,2,7,1,3]",
			SearchVal:    5,
			ExpectedTree: "[]",
		},
		{
			Name:         "Search: null tree",
			SourceTree:   "[]",
			SearchVal:    5,
			ExpectedTree: "[]",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			root, err := StringToIntTree(tc.SourceTree)
			assert.NoError(t, err)

			var out string
			out, err = IntTreeToString(searchBST(root, tc.SearchVal))
			assert.NoError(t, err)
			assert.Equal(t, tc.ExpectedTree, out)
		})
	}
}
