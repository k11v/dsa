package main

// LeetCode

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	nums3 := make([][]int, 0, max(len(nums1), len(nums2)))

	i := 0
	j := 0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i][0] == nums2[j][0] {
			id := nums1[i][0]
			sum := nums1[i][1] + nums2[j][1]
			nums3 = append(nums3, []int{id, sum})
			i++
			j++
		} else if nums1[i][0] < nums2[j][0] {
			nums3 = append(nums3, []int{nums1[i][0], nums1[i][1]})
			i++
		} else {
			nums3 = append(nums3, []int{nums2[j][0], nums2[j][1]})
			j++
		}
	}

	for i < len(nums1) {
		nums3 = append(nums3, []int{nums1[i][0], nums1[i][1]})
		i++
	}

	for j < len(nums2) {
		nums3 = append(nums3, []int{nums2[j][0], nums2[j][1]})
		j++
	}

	return nums3
}
