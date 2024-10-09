package main

// LeetCode

func minWindow(s string, t string) string {
	gotIndices := make([]int, 0)
	gotIndicesFromByte := make(map[byte][]int)
	wantFromByte := make(map[byte]int)
	incomplete := make(map[byte]struct{})
	for i := 0; i < len(t); i++ {
		gotIndicesFromByte[t[i]] = make([]int, 0)
		wantFromByte[t[i]]++
		incomplete[t[i]] = struct{}{}
	}

	findWant := func(p *int) (ok bool) {
		*p++
		for *p < len(s) && wantFromByte[s[*p]] == 0 {
			*p++
		}
		if *p == len(s) {
			return false
		}
		q := *p

		gotIndices = append(gotIndices, q)
		gotIndicesFromByte[s[q]] = append(gotIndicesFromByte[s[q]], q)
		if got, want := len(gotIndicesFromByte[s[q]]), wantFromByte[s[q]]; got == want {
			delete(incomplete, s[q])
		}

		return true
	}

	l := -1
	if !findWant(&l) {
		return ""
	}
	r := l

	minB, minE := 0, len(s)
	for {
		if len(incomplete) == 0 {
			b, e := l, r+1
			if e-b < minE-minB {
				minB, minE = b, e
			}
		}
		if !findWant(&r) {
			break
		}
		for len(gotIndicesFromByte[s[l]]) > wantFromByte[s[l]] {
			gotIndices = gotIndices[1:]
			gotIndicesFromByte[s[l]] = gotIndicesFromByte[s[l]][1:]
			l = gotIndices[0]
		}
	}

	if len(incomplete) != 0 {
		return ""
	}
	return s[minB:minE]
}
