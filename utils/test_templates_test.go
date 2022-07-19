package utils

import (
	"fmt"
	"testing"
)

func TestOneFunc(t *testing.T) {
	type TestCase struct {
		Name string
		//...
	}

	tcs := []TestCase{}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {

		})
	}
}

func TestMultipleFunc(t *testing.T) {
	fnNames := map[string]func( /*func arguments*/ ) /*returned value*/ {}

	type TestCase struct {
		Name string
		//...
	}

	tcs := []TestCase{}

	for fnName, fn := range fnNames {
		for _, tc := range tcs {
			t.Run(fmt.Sprintf("%s/%s", fnName, tc.Name), func(t *testing.T) {
				fn( /*tc.Arguments*/ )

				//	assert.Equal(t, )
			})
		}
	}
}
