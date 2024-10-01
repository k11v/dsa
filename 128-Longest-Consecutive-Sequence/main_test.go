package main

// LeetCode

func longestConsecutive(nums []int) int {
	longest := 0

	presentNumbers := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		presentNumbers[nums[i]] = struct{}{}
	}

	for i := 0; i < len(nums); i++ {
		if _, previousPresent := presentNumbers[nums[i]-1]; previousPresent {
			continue
		}

		start := nums[i]
		current := start
		for {
			if _, nextPresent := presentNumbers[current+1]; nextPresent {
				current = current+1
				continue
			}
			break
		}

		longest = max(longest, current-start+1)
	}

	return longest
}
