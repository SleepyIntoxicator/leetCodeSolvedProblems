package utils

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AssertEqualLists asserts that the specified expected(array, slice, string...) is equal to specified
// actual(array, slice, string).
func AssertEqualLists(t *testing.T, expected, actual interface{}) bool {
	if !isList(t, expected) || !isList(t, actual) {
		//if reflect.ValueOf(expected).Kind() == reflect.String &&
		//	reflect.ValueOf(actual).Kind() == reflect.String {
		//	expected = expected.([]byte)
		//	actual = actual.([]byte)
		//} else {
		//	return false
		//}
		return false
	}

	aValue := reflect.ValueOf(expected)
	bValue := reflect.ValueOf(actual)

	if isListOrMapEmpty(expected) && isListOrMapEmpty(actual) {
		return true
	}

	expectedLen := aValue.Len()
	actualLen := bValue.Len()

	if expectedLen != actualLen {
		var msg bytes.Buffer

		msg.WriteString("elements differ")
		msg.WriteString(fmt.Sprintf("\nexpected:\t%v", aValue))
		msg.WriteString(fmt.Sprintf("\nactual:\t\t%v", bValue))
		assert.Fail(t, msg.String())
		return false
	}

	var isIntSlice bool
	if expectedLen > 0 {
		//TODO: compare with aValue.Type().Elem().String() == "int"
		isIntSlice = reflect.ValueOf(aValue.Index(0).Interface()).Kind() == reflect.Int
	}

	var diff []byte
	if isIntSlice {
		diff = make([]byte, 0, expectedLen)
	}

	areEqual := true
	for i := 0; i < expectedLen; i++ {
		expected := aValue.Index(i).Interface()
		actual := bValue.Index(i).Interface()

		if !assert.ObjectsAreEqual(expected, actual) {
			areEqual = false
			if isIntSlice {
				diff = append(diff, '-')
			}
		} else if isIntSlice {
			diff = append(diff, ' ')
		}
	}

	if !areEqual {
		var msg bytes.Buffer

		msg.WriteString("elements differ")
		msg.WriteString(fmt.Sprintf("\nexpected:\t%v", aValue))
		if isIntSlice {
			msg.WriteString(fmt.Sprintf("\ndiff:\t\t%c", diff))
		}
		msg.WriteString(fmt.Sprintf("\nactual:\t\t%v", bValue))
		assert.Fail(t, msg.String())
		return false
	}

	return true
}

// isList checks that the provided value is array, slice or string.
func isList(t *testing.T, list interface{}) (ok bool) {
	kind := reflect.TypeOf(list).Kind()
	switch kind {
	case reflect.Array, reflect.Slice, reflect.String:
		return true
	case reflect.Map:
		assert.Fail(t, fmt.Sprintf("%v has unsupported type '%s', expecting array, slice or string", list, kind))
		return false
	default:
		assert.Fail(t, fmt.Sprintf("%q has unsupported type '%s', expecting array, slice or string", list, kind))
	}

	return false
}

// isListOrMapEmpty gets whether the specified list or map is empty or not.
// Returns false for all other types of object.
func isListOrMapEmpty(object interface{}) bool {
	objValue := reflect.ValueOf(object)

	switch objValue.Kind() {
	case reflect.Slice, reflect.Array, reflect.String, reflect.Map:
		return objValue.Len() == 0
	default:
		return false
	}
}
