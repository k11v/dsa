package main

// LeetCode

func singleNumber(nums []int) int {
	xor := 0
	for i, num := range nums {
		if i == 0 {
			xor = num
			continue
		}
		xor = xor ^ num
	}
	return xor
}
