package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
		{[]int{5, 25, 75}, 100, []int{2, 3}},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, want := twoSum(tt.numbers, tt.target), tt.want; !reflect.DeepEqual(got, want) {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

// LeetCode

func twoSum(numbers []int, target int) []int {
	i := 0
	j := len(numbers) - 1

	for j > i {
		gotSum := numbers[i] + numbers[j]
		wantSum := target

		switch {
		case gotSum == wantSum:
			return []int{i + 1, j + 1}
		case gotSum > wantSum:
			j--
		case gotSum < wantSum:
			i++
		default:
			panic("")
		}
	}

	panic("")
}
