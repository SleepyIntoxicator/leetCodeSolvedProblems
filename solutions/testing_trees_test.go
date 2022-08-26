package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToIntTree(t *testing.T) {
	type TestCase struct {
		Name             string
		SourceTreeValues string
		IsValid          bool
		ExpectedError    error
	}

	tcs := []TestCase{
		{
			Name:             "Valid Convert",
			SourceTreeValues: "[1,2,3]",
			IsValid:          true,
		},
		{
			Name:             "Valid Convert 2",
			SourceTreeValues: "[1,2,2,null,3]",
			IsValid:          true,
		},
		{
			Name:             "Valid Convert 3",
			SourceTreeValues: "[1,2,3,4,5,6,7]",
			IsValid:          true,
		},
		{
			Name:             "Valid Convert 4",
			SourceTreeValues: "[1,2,3,null,4,null,5,6,7,8,9,10]",
			IsValid:          true,
		},
		{
			Name:             "Valid Convert Of One Elem",
			SourceTreeValues: "[1]",
			IsValid:          true,
		},
		{
			Name:             "Valid Convert With Not Full Level",
			SourceTreeValues: "[1,2,3,4,null,null]",
			IsValid:          true,
		},
		{
			Name:             "Valid Empty Tree",
			SourceTreeValues: "[]",
			IsValid:          true,
		},
		{
			Name:             "Invalid Source data: empty string",
			SourceTreeValues: "",
			IsValid:          false,
			ExpectedError:    errEmptySourceString,
		},
		{
			Name:             "Invalid source data: empty tree",
			SourceTreeValues: "[null,null]",
			IsValid:          false,
			ExpectedError:    errInvalidSourceData,
		},
		{
			Name:             "Invalid source data: child elem torn off",
			SourceTreeValues: "[1,null,null,2]",
			IsValid:          false,
			ExpectedError:    errTornSourceData,
		},
		{
			Name:             "Invalid Convert With Full 'nil' Level",
			SourceTreeValues: "[1,2,3,4,null,null,null,null,null,5]",
			IsValid:          false,
			ExpectedError:    errTornSourceData,
		},
		{
			Name:             "Invalid Source Data",
			SourceTreeValues: "[1,2,3,BULL]",
			IsValid:          false,
			ExpectedError:    errInvalidSourceData,
		},
		{
			Name:             "Invalid EmptyItem Data",
			SourceTreeValues: "[1,2,3,]",
			IsValid:          false,
			ExpectedError:    errInvalidSourceData,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			root, err := StringToIntTree(tc.SourceTreeValues)
			if tc.IsValid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.ErrorIs(t, err, tc.ExpectedError)
				return
			}

			clearedTreeValues := clearEmptyTreeValuesFromEndOfString(tc.SourceTreeValues)

			var out string
			out, err = IntTreeToString(root)
			assert.NoError(t, err)
			assert.Equal(t, clearedTreeValues, out)
		})
	}
}

func TestIntTreeToString(t *testing.T) {
	type TestCase struct {
		Name             string
		SourceTreeValues string
	}

	tcs := []TestCase{
		{
			Name:             "Valid Convert",
			SourceTreeValues: "[1,2,3]",
		},
		{
			Name:             "Valid Convert 2",
			SourceTreeValues: "[1,2,3,4,5,6,7]",
		},
		{
			Name:             "Valid Convert 3",
			SourceTreeValues: "[1,2,3,null,4,null,5,6,7,8,9,10]",
		},
		{
			Name:             "Valid Convert Of One Elem",
			SourceTreeValues: "[1]",
		},
		{
			Name:             "Valid Empty Tree",
			SourceTreeValues: "[]",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			root, err := StringToIntTree(tc.SourceTreeValues)
			assert.NoError(t, err)

			var out string
			out, err = IntTreeToString(root)
			assert.NoError(t, err)
			assert.Equal(t, tc.SourceTreeValues, out)
		})
	}
}
