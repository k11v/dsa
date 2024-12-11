package main

import (
	"fmt"
	"io"
	"strconv"
)

func main() {
	countFromNumAndBlinks := make(map[struct{ num, blinks int }]int)

	var count func(num, blinks int) int
	count = func(num, blinks int) int {
		if countFromNumAndBlinks[struct{ num, blinks int }{num, blinks}] != 0 {
			return countFromNumAndBlinks[struct{ num, blinks int }{num, blinks}]
		}
		if blinks == 0 {
			return 1
		}
		var c int
		if num == 0 {
			c = count(1, blinks-1)
		} else if left, right, ok := splitIfEven(num); ok {
			c = count(left, blinks-1) + count(right, blinks-1)
		} else {
			c = count(num*2024, blinks-1)
		}
		countFromNumAndBlinks[struct{ num, blinks int }{num, blinks}] = c
		return c
	}

	answer := 0
	for {
		var num int
		_, err := fmt.Scanf("%d", &num)
		if err == io.EOF {
			break
		}
		answer += count(num, 25)
	}
	fmt.Println(answer)
}

func splitIfEven(num int) (left int, right int, ok bool) {
	numString := strconv.Itoa(num)
	if len(numString)%2 == 0 {
		m := len(numString) / 2
		left, err := strconv.Atoi(numString[:m])
		if err != nil {
			panic(err)
		}
		right, err := strconv.Atoi(numString[m:])
		if err != nil {
			panic(err)
		}
		return left, right, true
	}
	return 0, 0, false
}
