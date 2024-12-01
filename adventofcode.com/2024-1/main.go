package main

import (
	"io"
	"fmt"
	"slices"
	"flag"
)

func main() {
	flag.Parse()
	
	switch part := flag.Arg(0); part {
	case "1":
		part1()
	case "2":
		part2()
	default:
		panic(fmt.Errorf("unknown part %q", part))
	}
}

func part1() {
	left := make([]int, 0)
	right := make([]int, 0)

	var err error
	var l, r int
	for {
		_, err = fmt.Scanf("%d %d\n", &l, &r)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
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
}

func part2() {
	left := make(map[int]int, 0)
	right := make(map[int]int, 0)

	var err error
	var l, r int
	for {
		_, err = fmt.Scanf("%d %d\n", &l, &r)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
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

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}