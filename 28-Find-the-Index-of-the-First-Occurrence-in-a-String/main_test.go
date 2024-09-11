package main

import (
	"strconv"
	"testing"
)

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"sadbutsad", "sad", 0},
		{"leetcode", "leeto", -1},
		{"notsadnotsadatall", "sad", 3},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, want := strStr(tt.haystack, tt.needle), tt.want; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

// LeetCode

func prefixFunction(s string) []int {
	p := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		j := p[i-1] - 1

		for j != -1 && s[i] != s[j+1] {
			j = p[j] - 1
		}

		if s[i] == s[j+1] {
			p[i] = j + 2
		} else {
			p[i] = 0
		}
	}
	return p
}

func strStr(haystack string, needle string) int {
	p := prefixFunction(needle + "#" + haystack)
	for i, n := range p {
		if n == len(needle) {
			r := i
			r = r - n - 1 // remove (needle + "#") from equation
			r = r - n + 1 // adjust to return the first index instead of the last
			return r
		}
	}
	return -1
}
