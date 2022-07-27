package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertEqualSlices(t *testing.T) {
	type TestCase struct {
		Name           string
		Expected       interface{}
		Actual         interface{}
		ExpectedResult bool
	}

	tcs := []TestCase{
		{
			Name:           "Equal int slices",
			Expected:       []int{1, 2, 3},
			Actual:         []int{1, 2, 3},
			ExpectedResult: true,
		}, {
			Name:           "Empty int slices",
			Expected:       []int{},
			Actual:         []int{},
			ExpectedResult: true,
		},
		{
			Name:           "Diff length in int slices",
			Expected:       []int{1, 5, 6, 7, 3, 4, 5},
			Actual:         []int{1, 2, 3},
			ExpectedResult: false,
		},
		{
			Name:           "Diff length in int slices",
			Expected:       []int{1, 2, 3},
			Actual:         []int{1, 5, 6, 7, 3, 4, 5},
			ExpectedResult: false,
		},
		{
			Name:           "Not equal int slices",
			Expected:       []int{1, 5, 6, 7, 3, 4, 5},
			Actual:         []int{1, 0, 0, 2, 3, 4, 5},
			ExpectedResult: false,
		},
		{
			Name:           "Equal float32 slices",
			Expected:       []float32{1.25, 2.45, 3.7},
			Actual:         []float32{1.25, 2.45, 3.7},
			ExpectedResult: true,
		},
		{
			Name:           "Diff float32 slices",
			Expected:       []float32{1.25, 2.45, 3.7, 5.1},
			Actual:         []float32{1.25, 2.2, 3.0, 4.0},
			ExpectedResult: false,
		},
		{
			Name:           "Equal float64 slices",
			Expected:       []float64{1.25, 2.45, 3.7},
			Actual:         []float64{1.25, 2.45, 3.7},
			ExpectedResult: true,
		},
		{
			Name:           "Diff float64 slices",
			Expected:       []float64{1.25, 2.45, 3.7, 5.1},
			Actual:         []float64{1.25, 2.2, 3.0, 4.0},
			ExpectedResult: false,
		},
		{
			Name:           "Equal bool slices",
			Expected:       []bool{true, true, false},
			Actual:         []bool{true, true, false},
			ExpectedResult: true,
		},
		{
			Name:           "Diff bool slices",
			Expected:       []bool{true, true, false},
			Actual:         []bool{true, false, false},
			ExpectedResult: false,
		},
		{
			Name:           "Equal string slices",
			Expected:       []string{"aaa", "bbb", "ccc"},
			Actual:         []string{"aaa", "bbb", "ccc"},
			ExpectedResult: true,
		},
		{
			Name:           "Not equal string slices",
			Expected:       []string{"aaa", "bbb", "ccc"},
			Actual:         []string{"aaa", "ccc", "bbb"},
			ExpectedResult: false,
		},
		{
			Name:           "Equal types of the slice([]interface{}) elements",
			Expected:       []interface{}{1, "ccc", 3},
			Actual:         []interface{}{1, "ccc", 3},
			ExpectedResult: true,
		},
		{
			Name:           "Different types of the slice([]interface{}) elements",
			Expected:       []interface{}{1, 2, 3},
			Actual:         []interface{}{1, "ccc", 3},
			ExpectedResult: false,
		},
		{
			Name:           "Equal [int]int maps",
			Expected:       map[int]int{1: 1, 2: 2, 3: 3},
			Actual:         map[int]int{1: 1, 2: 2, 3: 3},
			ExpectedResult: false,
		},
		{
			Name:           "Diff [int]int maps",
			Expected:       map[int]int{1: 1, 2: 2, 3: 3},
			Actual:         map[int]int{1: 1, 2: 5, 3: 7},
			ExpectedResult: false,
		},
		{
			Name:           "Diff [int]int maps with diff lengths",
			Expected:       map[int]int{1: 1, 2: 2, 3: 3},
			Actual:         map[int]int{1: 1},
			ExpectedResult: false,
		},
		{
			Name:           "Success with not slices - string",
			Expected:       "string",
			Actual:         "string",
			ExpectedResult: true,
		},
		{
			Name:           "Fail with not slices - string",
			Expected:       "string",
			Actual:         "strong",
			ExpectedResult: false,
		},
		{
			Name:           "Not slices - int",
			Expected:       513,
			Actual:         513,
			ExpectedResult: false,
		},
	}

	tt := &testing.T{}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			isEqual := AssertEqualLists(tt, tc.Expected, tc.Actual)

			assert.Equal(t, tc.ExpectedResult, isEqual)
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type TestCase struct {
		Name     string
		Object   interface{}
		Expected bool
	}

	tcs := []TestCase{
		{
			Name:     "Empty int slice",
			Object:   []int{},
			Expected: true,
		},
		{
			Name:     "Not empty int slice",
			Object:   []int{1, 2, 3},
			Expected: false,
		},
		{
			Name:     "Empty int array with nulls",
			Object:   [3]int{},
			Expected: false,
		},
		{
			Name:     "Not empty int array",
			Object:   [3]int{1, 2, 3},
			Expected: false,
		},
		{
			Name:     "Empty string",
			Object:   "",
			Expected: true,
		},
		{
			Name:     "Not empty string",
			Object:   "string",
			Expected: false,
		},
		{
			Name:     "Nil map",
			Object:   map[int]int{},
			Expected: true,
		},
		{
			Name:     "Empty map",
			Object:   make(map[int]int, 0),
			Expected: true,
		},
		{
			Name:     "Not empty map",
			Object:   map[int]int{1: 1},
			Expected: false,
		},
		{
			Name:     "Not list",
			Object:   44,
			Expected: false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			out := isListOrMapEmpty(tc.Object)

			assert.Equal(t, tc.Expected, out)
		})
	}
}
