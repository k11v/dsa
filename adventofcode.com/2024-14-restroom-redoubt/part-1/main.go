package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var q1, q2, q3, q4 int
	n, m := 103, 101
	t := 100
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var px, py, vx, vy int
		_, err := fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		if err != nil {
			panic(err)
		}
		fx := addUntilNonNegative(px+t*vx, m) % m
		fy := addUntilNonNegative(py+t*vy, n) % n

		switch {
		case fx > m/2 && fy > n/2:
			q1++
		case fx > m/2 && fy < n/2:
			q2++
		case fx < m/2 && fy < n/2:
			q3++
		case fx < m/2 && fy > n/2:
			q4++
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(q1 * q2 * q3 * q4)
}

func addUntilNonNegative(c, o int) int {
	if c >= 0 {
		return c
	}
	m := abs(c) / o
	if abs(c)%o != 0 {
		m++
	}
	return c + m*o
}

func abs(v int) int {
	if v >= 0 {
		return v
	} else {
		return -v
	}
}
