package solutions

import "sort"

func containsDuplicateWithHashMap(nums []int) bool {
	duplicates := make(map[int]struct{}, 0)

	for i := range nums {
		_, ok := duplicates[nums[i]]
		if ok {
			return true
		}
		duplicates[nums[i]] = struct{}{}
	}
	return false
}

func containsDuplicateWithSort(nums []int) bool {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}
	return false
}
