package main

import (
	"errors"
	"fmt"
	"io"
)

func main() {
	next := func(x int) int {
		x = ((x * 64) ^ x) % 16777216
		x = ((x / 32) ^ x) % 16777216
		x = ((x * 2048) ^ x) % 16777216
		return x
	}

	s := 0
	for {
		var x int
		_, err := fmt.Scanf("%d", &x)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		for range 2000 {
			x = next(x)
		}
		s += x
	}
	fmt.Println(s)
}
