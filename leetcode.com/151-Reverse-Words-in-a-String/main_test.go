package main

import (
	"strconv"
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

func reverse(rs []rune, l, r int) {
	for i := 0; i < (r - l + 1) / 2; i++ {
		rs[l + i], rs[r - i] = rs[r - i], rs[l + i]
	}
}

func reverseWords(s string) string {
	rs := []rune(s)

	reverse(rs, 0, len(rs) - 1)
	wordStart := -1
	for i := 0; i < len(rs); i++ {
		if rs[i] == ' ' {
			if wordStart != -1 {
				reverse(rs, wordStart, i - 1)
				wordStart = -1
			}
		} else {
			if wordStart == -1 {
				wordStart = i
			}
		}
	}
	if wordStart != -1 {
		reverse(rs, wordStart, len(rs) - 1)
		wordStart = -1
	}

	w := 0
	for r := 0; r < len(rs); r++ {
		if rs[r] != ' ' {
			rs[w] = rs[r]
			w++
		} else {
			if w != 0 && r != len(rs) - 1 && rs[r + 1] != ' ' {
				rs[w] = ' '
				w++
			}
		}
	}
	rs = rs[0:w]

	return string(rs)
}
