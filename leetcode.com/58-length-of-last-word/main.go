package main

func lengthOfLastWord(s string) int {
	i := len(s) - 1

	for s[i] == ' ' {
		i--
	}
	end := i

	for i >= 0 && s[i] != ' ' {
		i--
	}
	start := i + 1

	return end - start + 1
}

