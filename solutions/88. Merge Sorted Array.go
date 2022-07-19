package solutions

func merge(nums1 []int, m int, nums2 []int, n int) {
	last1, last2, ins1Place := m-1, n-1, len(nums1)-1

	for last1 >= 0 && last2 >= 0 {
		if nums1[last1] > nums2[last2] {
			nums1[ins1Place] = nums1[last1]
			last1--
		} else if nums2[last2] >= nums1[last1] {
			nums1[ins1Place] = nums2[last2]
			last2--
		}
		ins1Place--
	}

	for last2 >= 0 {
		nums1[ins1Place] = nums2[last2]
		ins1Place--
		last2--
	}
}
