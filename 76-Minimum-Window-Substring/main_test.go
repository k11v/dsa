package main

// LeetCode

func minWindow(s string, t string) string {
	got := make(map[byte]int)
	want := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		want[t[i]]++
	}
	gotComplete := 0
	wantComplete := len(want)


	minB, minE := 0, len(s)
	b, e := 0, 0
	for e < len(s) {
		if want[s[e]] > 0 {
			got[s[e]]++
			if got[s[e]] == want[s[e]] {
				gotComplete++
			}
		}
		e++

		if gotComplete == wantComplete {
			for b < len(s) {
				if want[s[b]] > 0 {
					if got[s[b]] == want[s[b]] {
						break
					}
					got[s[b]]--
				}
				b++
			}

			if e-b < minE-minB {
				minB, minE = b, e
			}
		}
	}

	if gotComplete != wantComplete {
		return ""
	}
	return s[minB:minE]
}
