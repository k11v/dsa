package main

import (
	"slices"
	"strconv"
	"strings"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, want := reverseWords(tt.s), tt.want; got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		})
	}
}

// LeetCode

func reverseWords(s string) string {
	words := strings.Fields(s)
	slices.Reverse(words)
	return strings.Join(words, " ")
}
