package main

// LeetCode

func isValid(s string) bool {
	expected := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			expected = append(expected, ')')
		case '[':
			expected = append(expected, ']')
		case '{':
			expected = append(expected, '}')
		case ')', ']', '}':
			if len(expected) == 0 || expected[len(expected)-1] != s[i] {
				return false
			}
			expected = expected[:len(expected)-1]
		default:
			panic("unexpected input")
		}
	}
	return len(expected) == 0
}
