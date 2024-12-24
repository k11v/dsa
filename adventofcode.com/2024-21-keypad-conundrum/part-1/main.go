package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	type ij struct{ i, j int }
	keypadNumeric := map[rune]ij{
		'7': {0, 0},
		'8': {0, 1},
		'9': {0, 2},
		'4': {1, 0},
		'5': {1, 1},
		'6': {1, 2},
		'1': {2, 0},
		'2': {2, 1},
		'3': {2, 2},
		' ': {3, 0},
		'0': {3, 1},
		'A': {3, 2},
	}
	keypadDirectional := map[rune]ij{
		' ': {0, 0},
		'^': {0, 1},
		'A': {0, 2},
		'<': {1, 0},
		'v': {1, 1},
		'>': {1, 2},
	}

	pathNumeric := func(start rune, end rune) []rune {
		path := make([]rune, 0)
		si, sj := keypadNumeric[start].i, keypadNumeric[start].j
		ei, ej := keypadNumeric[end].i, keypadNumeric[end].j

		if si > ei {
			path = append(path, slices.Repeat([]rune{'^'}, si-ei)...)
			if sj > ej {
				path = append(path, slices.Repeat([]rune{'<'}, sj-ej)...)
			} else {
				path = append(path, slices.Repeat([]rune{'>'}, ej-sj)...)
			}
		} else {
			if sj > ej {
				path = append(path, slices.Repeat([]rune{'<'}, sj-ej)...)
			} else {
				path = append(path, slices.Repeat([]rune{'>'}, ej-sj)...)
			}
			path = append(path, slices.Repeat([]rune{'v'}, ei-si)...)
		}

		path = append(path, 'A')
		return path
	}

	pathDirectional := func(start rune, end rune) []rune {
		path := make([]rune, 0)
		si, sj := keypadDirectional[start].i, keypadDirectional[start].j
		ei, ej := keypadDirectional[end].i, keypadDirectional[end].j

		if si < ei {
			path = append(path, slices.Repeat([]rune{'v'}, ei-si)...)
			if sj > ej {
				path = append(path, slices.Repeat([]rune{'<'}, sj-ej)...)
			} else {
				path = append(path, slices.Repeat([]rune{'>'}, ej-sj)...)
			}
		} else {
			if sj > ej {
				path = append(path, slices.Repeat([]rune{'<'}, sj-ej)...)
			} else {
				path = append(path, slices.Repeat([]rune{'>'}, ej-sj)...)
			}
			path = append(path, slices.Repeat([]rune{'^'}, si-ei)...)
		}

		path = append(path, 'A')
		return path
	}

	answer := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		code := []rune(scanner.Text())

		path := make([]rune, 0)
		curr1 := 'A'
		for _, c1 := range code {
			curr2 := 'A'
			for _, c2 := range pathNumeric(curr1, c1) {
				curr3 := 'A'
				for _, c3 := range pathDirectional(curr2, c2) {
					path = append(path, pathDirectional(curr3, c3)...)
					curr3 = c3
				}
				curr2 = c2
			}
			curr1 = c1
		}

		var codeNumber int
		_, err := fmt.Sscanf(string(code), "%d", &codeNumber)
		if err != nil {
			panic(err)
		}

		fmt.Println(len(path), string(path), codeNumber)
		answer += len(path) * codeNumber
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(answer)
}
