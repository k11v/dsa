package main

// LeetCode

func minWindow(s string, t string) string {
	got := make(map[byte]int)
	want := make(map[byte]int)

	for i := 0; i < len(t); i++ {
		want[t[i]]++
	}

	l := 0
	for l < len(s) && want[s[l]] == 0 {
		l++
	}
	if l == len(s) {
		return ""
	}

	r := l
	gotComplete := 0
	wantComplete := len(want)
	for gotComplete < wantComplete {
		for r < len(s) && want[s[r]] == 0 {
			r++
		}
		if r == len(s) {
			return ""
		}
		got[s[r]]++
		if got[s[r]] == want[s[r]] {
			gotComplete++
		}
		r++
	}
	r--

	minL, minR := l, r
outer:
	for {
		for got[s[l]] > want[s[l]] {
			got[s[l]]--
			l++
			for l < len(s) && want[s[l]] == 0 {
				l++
			}
			if l == len(s) {
				break outer
			}
		}

		if r-l < minR-minL {
			minL, minR = l, r
		}

		r++
		for r < len(s) && want[s[r]] == 0 {
			r++
		}
		if r == len(s) {
			break outer
		}
		got[s[r]]++
	}

	return s[minL:minR+1]
}
