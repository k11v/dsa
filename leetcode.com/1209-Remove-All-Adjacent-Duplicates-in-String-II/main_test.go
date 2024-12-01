package main

import (
	"strings"
)

// LeetCode

func removeDuplicates(s string, k int) string {
	runes := make([]rune, 0)
	counts := make([]int, 0)

	for _, r := range s {
		n := len(runes)
		if n != 0 && runes[n-1] == r {
			counts[n-1] += 1
			if counts[n-1] == k {
				runes = runes[:n-1]
				counts = counts[:n-1]
			}
		} else {
			runes = append(runes, r)
			counts = append(counts, 1)
		}
	}

	w := strings.Builder{}
	for i := 0; i < len(runes); i++ {
		w.WriteString(strings.Repeat(string(runes[i]), counts[i]))
	}
	return w.String()
}
