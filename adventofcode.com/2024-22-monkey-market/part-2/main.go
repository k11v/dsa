package main

import (
	"errors"
	"fmt"
	"io"
	"reflect"
)

func main() {
	next := func(x int) int {
		x = ((x * 64) ^ x) % 16777216
		x = ((x / 32) ^ x) % 16777216
		x = ((x * 2048) ^ x) % 16777216
		return x
	}

	type abcd struct{ a, b, c, d int }
	countFromSeq := make(map[abcd]int)
	for buyerNumber := 0; ; buyerNumber++ {
		var x int
		_, err := fmt.Scanf("%d", &x)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}

		prevX := 0
		diffs := make([]int, 0)
		seqSeen := make(map[abcd]struct{})
		for range 2000 {
			prevX = x
			x = next(x)
			prevPrice := prevX % 10
			currPrice := x % 10
			diff := currPrice - prevPrice
			diffs = append(diffs, diff)
			if len(diffs) >= 4 {
				diffs = diffs[len(diffs)-4:]
				seq := abcd{diffs[0], diffs[1], diffs[2], diffs[3]}
				if reflect.DeepEqual(seq, abcd{-2, 1, -1, 3}) {
					fmt.Println(buyerNumber, currPrice)
				}
				if _, seen := seqSeen[seq]; !seen {
					countFromSeq[seq] += currPrice
					seqSeen[seq] = struct{}{}
				}
			}
		}
	}

	maxCount := -1
	for seq, count := range countFromSeq {
		if reflect.DeepEqual(seq, abcd{-2, 1, -1, 3}) {
			fmt.Println(seq, count)
		}
		maxCount = max(maxCount, count)
		if count == maxCount {
		}
	}

	fmt.Println(maxCount)
}
