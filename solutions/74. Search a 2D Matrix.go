package solutions

func searchMatrix(matrix [][]int, target int) bool {
	hl, hr, hm := 0, len(matrix)-1, 0

	for hl <= hr {
		hm = hl + (hr-hl)/2

		if target == matrix[hm][0] {
			return true
		}
		if target < matrix[hm][0] {
			hr = hm - 1
		} else if target > matrix[hm][0] {
			if hm == len(matrix)-1 || target < matrix[hm+1][0] {
				return binarySearch(matrix[hm], target)
			}
			hl = hm + 1
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	l, r, m := 0, len(nums)-1, 0
	for l <= r {
		m = l + (r-l)/2

		if target == nums[m] {
			return true
		} else if target < nums[m] {
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		}
	}
	return false
}
