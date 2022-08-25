package solutions

import (
	"errors"
	"strconv"
	"strings"
)

var (
	errEmptySourceString = errors.New("empty source string")
	errInvalidSourceData = errors.New("invalid source data")
	errTornSourceData    = errors.New("there are ghost elements in the source data")
)

// 7.5 hours to solve with tests

func clearEmptyTreeValuesFromEndOfString(treeValues string) string {
	treeValues = strings.TrimSuffix(treeValues, "]")

	for strings.HasSuffix(treeValues, "null") {
		treeValues = strings.TrimSuffix(treeValues, "null")
		if len(treeValues) > 1 {
			treeValues = strings.TrimSuffix(treeValues, ",")
		}
	}

	return treeValues + "]"
}

func convertStringSliceToIntSlice(src []string) ([]int, error) {
	intSlice := make([]int, len(src))
	for i, v := range src {
		if v == "null" {
			continue
		}

		var err error
		if intSlice[i], err = strconv.Atoi(src[i]); err != nil {
			return nil, errInvalidSourceData
		}
	}

	return intSlice, nil
}

// StringToIntTree fills the tree with elements from the source string.
// Returns errInvalidSourceData if the source contain invalid data.
// Returns errEmptySourceString if the source is empty string.
// Returns errTornSourceData if there are ghost elements in the tree.
func StringToIntTree(source string) (*TreeNode, error) {
	if source == "" {
		return nil, errEmptySourceString
	}

	// Exit if the tree is null
	if source == "[]" {
		return nil, nil
	}

	source = clearEmptyTreeValuesFromEndOfString(source)

	source = strings.TrimPrefix(source, "[")
	source = strings.TrimSuffix(source, "]")
	values := strings.Split(source, ",")

	if len(values) == 1 && values[0] == "" {
		return nil, errInvalidSourceData
	}

	intValues, err := convertStringSliceToIntSlice(values)
	if err != nil {
		return nil, err
	}

	// Fill in a slice of elements
	treeNodes := make([]*TreeNode, 0, len(values))
	for i, v := range values {
		if v == "null" {
			treeNodes = append(treeNodes, nil)
		} else {
			treeNodes = append(treeNodes, &TreeNode{Val: intValues[i]})
		}
	}

	root := treeNodes[0]

	prevLevelStartIndex, prevLevelEndIndex := 0, 1
	currIndex := prevLevelEndIndex

	for currIndex < len(treeNodes) {
		currLevelStart := currIndex
		countOfNullElems := 0

		// Filling in the current level of the tree
		for _, prev := range treeNodes[prevLevelStartIndex:prevLevelEndIndex] {
			if prev == nil {
				countOfNullElems++
				continue
			}

			prev.Left = treeNodes[currIndex]

			if currIndex+1 >= len(treeNodes) {
				currIndex += 1
				break
			}
			prev.Right = treeNodes[currIndex+1]

			currIndex += 2
		}

		// If all level elements are nil, an error is returned
		if prevLevelEndIndex-prevLevelStartIndex == countOfNullElems {
			return nil, errTornSourceData
		}

		// Preparing for the next level
		prevLevelStartIndex = currLevelStart
		prevLevelEndIndex = currIndex
	}

	return root, nil
}

func IntTreeToString(root *TreeNode) (string, error) {
	if root == nil {
		return "[]", nil
	}

	var valuesString []string

	// use BFS -
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		itemsOnLevel := len(queue)
		for _, item := range queue {
			if item == nil {
				valuesString = append(valuesString, "null")
			} else {
				valuesString = append(valuesString, strconv.Itoa(item.Val))
				queue = append(queue, item.Left, item.Right)
			}
		}

		queue = queue[itemsOnLevel:]
	}

	out := "[" + strings.Join(valuesString, ",") + "]"
	out = clearEmptyTreeValuesFromEndOfString(out)

	return out, nil
}

/*
        1
     L /|\ R
      /	| \
     2	|  3
  L / \R|L/ \ R
   4   5|6   7
  /___________\

*/
