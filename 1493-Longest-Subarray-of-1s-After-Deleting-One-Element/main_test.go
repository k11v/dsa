package main

// LeetCode

func longestSubarray(nums []int) int {
	maxOneCountWithJoin := 0
	prevOneCount := 0
	hasZeros := false

	for i := 0; i < len(nums); {
		currZeroCount := 0
		for i < len(nums) && nums[i] == 0 {
			hasZeros = true
			currZeroCount++
			i++
		}
		currOneCount := 0
		for i < len(nums) && nums[i] == 1 {
			currOneCount++
			i++
		}

		if currZeroCount == 1 {
			maxOneCountWithJoin = max(maxOneCountWithJoin, prevOneCount + currOneCount)
		} else {
			maxOneCountWithJoin = max(maxOneCountWithJoin, currOneCount)
		}

		prevOneCount = currOneCount
	}

	if !hasZeros {
		maxOneCountWithJoin--
	}

	return maxOneCountWithJoin
}
