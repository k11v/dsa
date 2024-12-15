package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
)

func main() {
	grid := make([][]rune, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		row := []rune(line)
		grid = append(grid, row)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	n, m := len(grid), len(grid[0])

	px, py := -1, -1
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if grid[x][y] == '@' {
				px, py = x, y
			}
		}
	}
	if px == -1 || py == -1 {
		panic("player not found")
	}

	var move func(x, y, mx, my int) (newX int, newY int, ok bool)
	move = func(x, y, mx, my int) (newX int, newY int, ok bool) {
		switch grid[x+mx][y+my] {
		case '#':
			return x, y, false
		case 'O':
			if _, _, ok := move(x+mx, y+my, mx, my); !ok {
				return x, y, false
			}
			fallthrough
		case '.':
			grid[x+mx][y+my] = grid[x][y]
			grid[x][y] = '.'
			return x + mx, y + my, true
		default:
			panic("bad cell")
		}
	}

	for m := range scanMovements(scanner) {
		var mx, my int
		switch m {
		case '^':
			mx, my = -1, 0
		case '>':
			mx, my = 0, 1
		case 'v':
			mx, my = 1, 0
		case '<':
			mx, my = 0, -1
		default:
			panic("bad movement")
		}
		if newPx, newPy, ok := move(px, py, mx, my); ok {
			px, py = newPx, newPy
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	checksum := 0
	for x := 0; x < n; x++ {
		for y := 0; y < m; y++ {
			if grid[x][y] == 'O' {
				checksum += x*100 + y
			}
		}
	}

	fmt.Println(checksum)
}

func scanMovements(scanner *bufio.Scanner) iter.Seq[rune] {
	return func(yield func(rune) bool) {
		for scanner.Scan() {
			for _, m := range scanner.Text() {
				if !yield(m) {
					return
				}
			}
		}
	}
}
