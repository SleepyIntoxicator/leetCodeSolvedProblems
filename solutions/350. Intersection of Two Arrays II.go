package solutions

func intersectBruteforce(nums1 []int, nums2 []int) []int {
	var short, long []int
	if len(nums1) < len(nums2) {
		short = nums1
		long = nums2
	} else {
		short = nums2
		long = nums1
	}

	shortHashtable := make(map[int]int, len(nums1))
	for _, v := range short {
		if _, ok := shortHashtable[v]; ok {
			shortHashtable[v]++
		} else {
			shortHashtable[v] = 1
		}
	}

	var intersects []int

	for _, v := range long {
		if count, ok := shortHashtable[v]; ok {
			intersects = append(intersects, v)
			if count == 1 {
				delete(shortHashtable, v)
			} else {
				shortHashtable[v]--
			}
		}
	}

	return intersects
}

func intersectWithSorting(nums1 []int, nums2 []int) []int {

	return nums1
}

func intersectWithBinarySearch(nums1 []int, nums2 []int) []int {

	return nums1
}

// What if the given array is already sorted? How would you optimize your algorithm?
func intersectAlreadySorted(nums1 []int, nums2 []int) []int {

	return nums1
}

// What if nums1's size is small compared to nums2's size? Which algorithm is better?
func intersectNums1SizeIsSmall(nums1 []int, nums2 []int) []int {

	return nums1
}

// What if elements of nums2 are stored on disk, and the memory is limited
// such that you cannot load all elements into the memory at once?
func intersectNums2OnDiskAndMassive(nums1 []int, nums2 []int) []int {

	return nums1
}
