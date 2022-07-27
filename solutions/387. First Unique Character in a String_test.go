package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstUniqChar(t *testing.T) {
	fnNames := map[string]func(s string) int{
		"firstUniqCharFromLeetCode": firstUniqCharFromLeetCode,
		"firstUniqCharFirst":        firstUniqCharFirst,
	}

	type TestCase struct {
		Name     string
		S        string
		Expected int
	}

	tcs := []TestCase{
		{
			Name:     "1",
			S:        "leetcode",
			Expected: 0,
		},
		{
			Name:     "2",
			S:        "loveleetcode",
			Expected: 2,
		},
		{
			Name:     "3",
			S:        "aabb",
			Expected: -1,
		},
	}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				assert.Equal(t, tc.Expected, fn(tc.S))
			})
		}
	}
}
