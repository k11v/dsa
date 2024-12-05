package main

import (
	"io"
	"fmt"
	"slices"
	"flag"
)

func main() {
	if err := run(); err != nil {
		panic()
	}
}

func run() error {
	left := make(map[int]int, 0)
	right := make(map[int]int, 0)

	var err error
	var l, r int
	for {
		_, err = fmt.Scanf("%d %d\n", &l, &r)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		left[l]++
		right[r]++
	}

	s := 0
	for x := range left {
		s += x * left[x] * right[x]
	}

	fmt.Println(s)
}
