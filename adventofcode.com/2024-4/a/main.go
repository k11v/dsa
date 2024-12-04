package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type ij struct{ i, j int }

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	grid := make([][]rune, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if len(grid) == 0 {
		return errors.New("empty grid")
	}

	count := 0

	n := len(grid)
	m := len(grid[0])
	chars := []rune{'X', 'M', 'A', 'S'}
	casesWithCharOffsets := [][]ij{
		{{+0, +0}, {+1, +0}, {+2, +0}, {+3, +0}},
		{{+0, +0}, {-1, +0}, {-2, +0}, {-3, +0}},
		{{+0, +0}, {+0, +1}, {+0, +2}, {+0, +3}},
		{{+0, +0}, {+0, -1}, {+0, -2}, {+0, -3}},
		{{+0, +0}, {+1, +1}, {+2, +2}, {+3, +3}},
		{{+0, +0}, {-1, -1}, {-2, -2}, {-3, -3}},
		{{+0, +0}, {+1, -1}, {+2, -2}, {+3, -3}},
		{{+0, +0}, {-1, +1}, {-2, +2}, {-3, +3}},
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
		loopJ:
			for _, offsets := range casesWithCharOffsets {
				for k, offset := range offsets {
					charI, charJ := i+offset.i, j+offset.j
					if charI < 0 || charI >= n || charJ < 0 || charJ >= m {
						continue loopJ
					}
					if grid[charI][charJ] != chars[k] {
						continue loopJ
					}
				}
				count++
			}
		}
	}

	fmt.Println(count)

	return nil
}
