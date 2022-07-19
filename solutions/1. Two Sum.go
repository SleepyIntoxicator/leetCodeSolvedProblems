package solutions

func twoSumNew(nums []int, target int) []int {
	elems := make(map[int]int, 0)
	for i, v := range nums {
		if it, ok := elems[v]; !ok {
			elems[target-v] = i
		} else {
			return []int{i, it}
		}
	}
	return []int{}
}

func twoSumOld(nums []int, target int) []int {
	hashMap := make(map[int]int, len(nums))

	for index := range nums {
		foundIndex, ok := hashMap[nums[index]]
		if ok {
			return []int{foundIndex, index}
		}

		ValueToFind := target - nums[index]
		hashMap[ValueToFind] = index
	}

	return nil
}

func twoSumBruteforce(nums []int, target int) []int {
	for i, v := range nums {
		for j, v2 := range nums {
			if i != j && (target-v == v2 || target-v2 == v) {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
