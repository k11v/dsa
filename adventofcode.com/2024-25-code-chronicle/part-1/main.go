package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	keys := make([][]int, 0)
	locks := make([][]int, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var typ string
		var heights []int
		switch scanner.Text() {
		case ".....":
			typ = "key"
			heights = []int{0, 0, 0, 0, 0}
		case "#####":
			typ = "lock"
			heights = []int{1, 1, 1, 1, 1}
		default:
			panic("invalid schematic")
		}

		for scanner.Scan() {
			if scanner.Text() == "" {
				break
			}
			for i, ch := range []rune(scanner.Text()) {
				switch ch {
				case '.':
				case '#':
					heights[i]++
				default:
					panic("invalid schematic")
				}
			}
		}
		for scanner.Err() != nil {
			panic(scanner.Err())
		}

		switch typ {
		case "key":
			keys = append(keys, heights)
		case "lock":
			locks = append(locks, heights)
		default:
			panic("invalid type")
		}
	}
	for scanner.Err() != nil {
		panic(scanner.Err())
	}
	n := len(keys[0])
	m := 7

	count := 0
	for _, key := range keys {
	lockLoop:
		for _, lock := range locks {
			for i := range n {
				fits := m-lock[i] >= key[i]
				if !fits {
					continue lockLoop
				}
			}
			count++
		}
	}

	fmt.Println(count)
}
