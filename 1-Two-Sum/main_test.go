package main

import (
	"reflect"
	"slices"
	"strconv"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{-3, 4, 3, 90}, 0, []int{0, 2}},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, want := twoSum(tt.nums, tt.target), tt.want; !reflect.DeepEqual(slices.Sorted(slices.Values(got)), want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

// LeetCode

func twoSum(nums []int, target int) []int {
	indexFromNum := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if j, ok := indexFromNum[target-nums[i]]; ok {
			return []int{i, j}
		}
		indexFromNum[nums[i]] = i
	}

	panic("unreachable")
}
