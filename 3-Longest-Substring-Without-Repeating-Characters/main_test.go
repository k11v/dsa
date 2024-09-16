package main

import (
	"strconv"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"aabcabcbb", 3},
		{"abba", 2},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, want := lengthOfLongestSubstring(tt.s), tt.want; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

// LeetCode

func lengthOfLongestSubstring(s string) int {
	indexFromCharacter := make(map[rune]int)
	startIndex := 0
	maxLength := 0

	for i, c := range s {
		if characterIndex, seen := indexFromCharacter[c]; seen {
			maxLength = max(maxLength, i-startIndex)
			startIndex = max(startIndex, characterIndex+1)
		}
		indexFromCharacter[c] = i
	}
	maxLength = max(maxLength, len(s)-startIndex)

	return maxLength
}
