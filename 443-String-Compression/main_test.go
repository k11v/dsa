package main

import "strconv"

// LeetCode

func compress(chars []byte) int {
	r, w := 0, 0
	for r < len(chars) {
		rn := r+1
		for rn < len(chars) && chars[rn] == chars[r] {
			rn++
		}

		chars[w] = chars[r]
		w++

		c := rn-r
		if c > 1 {
			cb := []byte(strconv.Itoa(c))
			for i := 0; i < len(cb); i++ {
				chars[w] = cb[i]
				w++
			}
		}

		r = rn
	}
	chars = chars[:w]
	return w
}
