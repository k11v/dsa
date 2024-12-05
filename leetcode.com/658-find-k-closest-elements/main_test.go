package main

import (
	"testing"
	"reflect"
)

func TestFindClosestElements(t *testing.T) {
	if got, want := findClosestElements([]int{1,1,1,10,10,10}, 1, 9), []int{10}; !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// LeetCode

func findClosestElements(arr []int, k int, x int) []int {
	l := 0
	r := len(arr)

	for r-l != 0 {
		c := (l+r-1) / 2
		if arr[c] == x {
			l, r = c, c+1
			break
		} else if arr[c] < x {
			l, r = c+1, r
		} else if arr[c] > x {
			l, r = l, c
		}
	}

	for i := r-l; i < k; i++ {
		if l-1 >= 0 && r+1 <= len(arr) {
			if abs(arr[l-1] - x) <= abs(arr[r] - x) {
				l = l-1
			} else {
				r = r+1
			}
		} else if l-1 >= 0 {
			l = l-1
		} else if r+1 <= len(arr) {
			r = r+1
		} else {
			break
		}
	}

	closestElements := make([]int, 0)
	for i := l; i < r; i++ {
		closestElements = append(closestElements, arr[i])
	}

	return closestElements
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
