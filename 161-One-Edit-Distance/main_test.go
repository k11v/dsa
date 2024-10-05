package main

import (
	"testing"
	"strconv"
)

func TestOneEditDistance(t *testing.T) {
	tests := []struct{
		s string
		t string
		want bool
	}{
		{"ab", "acb", true},
		{"cab", "ad", false},
		{"1203", "1213", true},
		{"acb", "ab", true},
		{"abc", "abc", false},
		{"", "", false},
		{"", "a", true},
		{"a", "", true},
		{"a", "b", true},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, want := oneEditDistance(tt.s, tt.t), tt.want; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

// LeetCode (like)

// oneEditDistance hasn't been tested on LeetCode, therefore could be incorrect.
func oneEditDistance(s, t string) bool {
	edits := 0

	i := 0
	j := 0
	for i < len(s) || j < len(t) {
		if i >= len(s) || j >= len(t) || s[i] != t[j] {
			if edits > 0 {
				return false
			}
			edits++

			if len(s) < len(t) {
				i--
			} else if len(s) > len(t) {
				j--
			} else {
				// No-op.
			}
		}

		i++
		j++
	}

	return i == len(s) && j == len(t) && edits == 1
}
