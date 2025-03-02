package main

func romanToInt(s string) int {
	v := 0
	p := 0

	for _, ch := range s {
		var d int
		switch ch {
		case 'I':
			d = 1
		case 'V':
			d = 5
		case 'X':
			d = 10
		case 'L':
			d = 50
		case 'C':
			d = 100
		case 'D':
			d = 500
		case 'M':
			d = 1000
		default:
			panic("unknown symbol")
		}

		v += d
		if d > p {
			v -= 2 * p
		}
		p = d
	}

	return v
}
