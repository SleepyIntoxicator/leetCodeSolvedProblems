package solutions

func maxSubArrayShort(nums []int) int {
	maxSubSum := -100000
	currSubsSum := 0

	for i := 0; i < len(nums); i++ {
		currSubsSum += nums[i]
		if currSubsSum > maxSubSum {
			maxSubSum = currSubsSum
		}
		if currSubsSum < 0 {
			currSubsSum = 0
		}
	}

	return maxSubSum
}

func maxSubArrayFirst(nums []int) int {
	maxSubSum := -100000
	currSubsSum := 0

	for i := 0; i < len(nums); i++ {

		if nums[i] >= 0 {
			currSubsSum += nums[i]
			if currSubsSum > maxSubSum {
				maxSubSum = currSubsSum
			}
		} else if nums[i] < 0 {
			if currSubsSum == 0 && nums[i] > maxSubSum {
				maxSubSum = nums[i]
			} else if currSubsSum+nums[i] < 0 {
				currSubsSum = 0
			} else if currSubsSum+nums[i] >= 0 {
				currSubsSum += nums[i]
			}
		}
	}

	return maxSubSum
}
