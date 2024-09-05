package main

import (
	"strconv"
	"testing"
	"unicode"
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
		l := unicode.ToLower(rune(s[i]))
		r := unicode.ToLower(rune(s[j]))

		switch {
		case !('a' <= l && l <= 'z' || '0' <= l && l <= '9'):
			i++
		case !('a' <= r && r <= 'z' || '0' <= r && r <= '9'):
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
