package main

import (
	"strconv"
	"strings"
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
	return strings.Index(haystack, needle)
}
