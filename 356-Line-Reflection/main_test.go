package main

import (
	"testing"
	"strconv"
)

func TestLineReflection(t *testing.T) {
	tests := []struct{
		points [][2]int
		want bool
	}{
		{[][2]int{[2]int{1, 1}, [2]int{-1, 1}}, true},
		{[][2]int{{1, 1}, {-1, -1}}, false},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, want := lineReflection(tt.points), tt.want; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

// LeetCode (like)

// lineReflection assumes that the reflection line cannot have fractional x.
// It also assumes that points do not repeat.
// It hasn't been tested with LeetCode.
func lineReflection(points [][2]int) bool {
	if len(points) == 0 {
		return true
	}

	s := 0
	c := 0
	for i := 0; i < len(points); i++ {
		if points[i][1] == points[0][1] {
			s += points[i][0]
			c++
		}
	}
	rx, remainder := s / c, s % c
	if remainder != 0 {
		return false
	}

	unbalanced := make(map[[2]int]struct{})
	for i := 0; i < len(points); i++ {
		px, py := points[i][0], points[i][1]

		var qx, qy int
		if px < rx {
			qx = px + 2*(rx-px)
		} else {
			qx = px - 2*(px-rx)
		}
		qy = py

		if px == rx {
			continue
		}
		if _, present := unbalanced[[2]int{qx, qy}]; present {
			delete(unbalanced, [2]int{qx, qy})
			continue
		}
		unbalanced[[2]int{px, py}] = struct{}{}
	}

	return len(unbalanced) == 0
}
