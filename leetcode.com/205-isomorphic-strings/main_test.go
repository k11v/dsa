package main

// LeetCode

func isIsomorphic(s string, t string) bool {
	alphabet := make(map[byte]byte)
	used := make(map[byte]struct{})

	for i := 0; i < len(s); i++ {
		if other, ok := alphabet[s[i]]; ok {
			if t[i] != other {
				return false
			}
			continue
		}
		if _, ok := used[t[i]]; ok {
			return false
		}
		alphabet[s[i]] = t[i]
		used[t[i]] = struct{}{}
	}

	return true
}
