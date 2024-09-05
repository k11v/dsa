package main

import (
	"strconv"
	"strings"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got, want := isPalindrome(tt.s), tt.want; got != want {
				t.Errorf("got %v, want %v", got, want)
			}
		})
	}
}

// LeetCode

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		l := strings.ToLower(string([]byte{s[i]}))[0]
		r := strings.ToLower(string([]byte{s[j]}))[0]

		switch {
		case !("a"[0] <= l && l <= "z"[0] || "0"[0] <= l && l <= "9"[0]):
			i++
		case !("a"[0] <= r && r <= "z"[0] || "0"[0] <= r && r <= "9"[0]):
			j--
		case l != r:
			return false
		case l == r:
			i++
			j--
		default:
			panic("")
		}
	}

	return true
}
