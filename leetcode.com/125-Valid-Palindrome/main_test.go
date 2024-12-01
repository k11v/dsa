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

		// The unicode package solution is not the simplest by any means.
		// It is slower than simple rune comparisons.
		// It is probably more useful in the real world though.
		switch {
		case !(unicode.IsLetter(l) || unicode.IsDigit(l)):
			i++
		case !(unicode.IsLetter(r) || unicode.IsDigit(r)):
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
