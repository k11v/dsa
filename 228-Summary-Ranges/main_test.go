package main

import (
	"fmt"
)

// LeetCode

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	ranges := make([]string, 0)
	start := nums[0]

	for i := 0; i < len(nums); i++ {
		if i != len(nums)-1 && nums[i+1] - nums[i] <= 1 {
			continue
		}

		end := nums[i]
		var r string
		if start == end {
			r = fmt.Sprintf("%d", start)
		} else {
			r = fmt.Sprintf("%d->%d", start, end)
		}
		ranges = append(ranges, r)

		if i != len(nums)-1 {
			start = nums[i+1]
		}
	}

	return ranges
}
