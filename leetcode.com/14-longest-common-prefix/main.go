package main

func longestCommonPrefix(strs []string) string {
	r := 0

Loop:
	for {
		for _, str := range strs {
			if r >= len(str) {
				break Loop
			}
			if str[r] != strs[0][r] {
				break Loop
			}
		}
		r++
	}

	return strs[0][:r]
}

