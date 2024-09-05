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

func strStr(haystack string, needle string) int {
outer:
	for i := 0; i < len(haystack); i++ {
		for j := 0; j < len(needle); j++ {
			if !(i+j < len(haystack) && haystack[i+j] == needle[j]) {
				continue outer
			}
		}
		return i
	}
	return -1
}
