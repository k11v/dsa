package main

import (
	"io"
	"fmt"
	"slices"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	left := make([]int, 0)
	right := make([]int, 0)

	var err error
	var l, r int
	for {
		_, err = fmt.Scanf("%d %d\n", &l, &r)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	s := 0
	for i := 0; i < len(left); i++ {
		s += abs(left[i] - right[i])
	}

	fmt.Println(s)
	return nil
}

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}
