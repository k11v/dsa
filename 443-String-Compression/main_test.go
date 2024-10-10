package main

import "strconv"

// LeetCode

func compress(chars []byte) int {
	w := 0
	c := 0

	flush := func() {
		if c > 1 {
			cb := []byte(strconv.Itoa(c))
			for i := 0; i < len(cb); i++ {
				chars[w] = cb[i]
				w++
			}
		}
	}

	for r := 0; r < len(chars); r++ {
		if r == 0 {
			w++
			c++
			continue
		}

		if chars[r] == chars[w-1] {
			c++
		} else {
			flush()
			c = 1
			chars[w] = chars[r]
			w++
		}
	}
	flush()

	return w
}
