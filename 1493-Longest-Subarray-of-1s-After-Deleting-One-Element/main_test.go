package main

// LeetCode

func longestSubarray(nums []int) int {
	maxOnes := 0

	anyZeros := false
	zeros := 0
	start := 0

	for end := 0; end < len(nums); end++ {
		if nums[end] == 0 {
			anyZeros = true
			zeros++
		}

		for zeros > 1 {
			if nums[start] == 0 {
				zeros--
			}
			start++
		}

		maxOnes = max(end-start+1 - zeros, maxOnes)
	}

	if !anyZeros {
		maxOnes--
	}

	return maxOnes
}
