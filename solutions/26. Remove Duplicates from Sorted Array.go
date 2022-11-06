package solutions

func RemoveDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	var offset int

	for i := 1; i < len(nums); i++ {
		if nums[i-offset-1] == nums[i] {
			offset++
		} else if offset > 0 {
			nums[i-offset] = nums[i]
		}
	}

	return len(nums) - offset
}

func RemoveDuplicatesFromLeetCode(nums []int) int {
	//Added by Slintox
	if len(nums) == 0 {
		return 0
	}
	//

	idx := 0
	for k := range nums {
		if nums[idx] != nums[k] {
			idx++
			nums[idx] = nums[k]
		}
	}
	return idx + 1
}
