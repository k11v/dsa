package main

// LeetCode

func minWindow(s string, t string) string {
	got, want := [128]int{}, [128]int{}
	gotComplete, wantComplete := 0, 0
	for i := 0; i < len(t); i++ {
		if want[t[i]] == 0 {
			wantComplete++
		}
		want[t[i]]++
	}

	found := false
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

		for gotComplete == wantComplete {
			found = true
			if e-b < minE-minB {
				minB, minE = b, e
			}

			if want[s[b]] > 0 {
				if got[s[b]] == want[s[b]] {
					gotComplete--
				}
				got[s[b]]--
			}
			b++
		}
	}

	if !found {
		return ""
	}
	return s[minB:minE]
}
