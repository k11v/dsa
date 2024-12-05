package main

// LeetCode

func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if s[i] != s[j] {
			return validPalindromeWithIgnoreIndex(s, i) || validPalindromeWithIgnoreIndex(s, j)
		}
		i++
		j--
	}

	return true
}

func validPalindromeWithIgnoreIndex(s string, ignoreIndex int) bool {
	i := 0
	j := len(s) - 1

	for i < j {
		if i == ignoreIndex {
			i++
		}
		if j == ignoreIndex {
			j--
		}

		if s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
