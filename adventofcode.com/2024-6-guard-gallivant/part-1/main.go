package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ma := make([][]rune, 0)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		ma = append(ma, []rune(line))
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	n, m := len(ma), len(ma[0])

	var gi, gj int
	var gd rune
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if ma[i][j] == '^' {
				gi, gj = i, j
				gd = 'U'
				ma[i][j] = 'X'
			}
		}
	}

loop:
	for {
		switch gd {
		case 'U':
			for gi-1 >= 0 && ma[gi-1][gj] != '#' {
				gi--
				ma[gi][gj] = 'X'
			}
			if gi-1 < 0 {
				break loop
			}
			gd = 'R'
		case 'R':
			for gj+1 < m && ma[gi][gj+1] != '#' {
				gj++
				ma[gi][gj] = 'X'
			}
			if gj+1 >= m {
				break loop
			}
			gd = 'D'
		case 'D':
			for gi+1 < n && ma[gi+1][gj] != '#' {
				gi++
				ma[gi][gj] = 'X'
			}
			if gi+1 >= n {
				break loop
			}
			gd = 'L'
		case 'L':
			for gj-1 >= 0 && ma[gi][gj-1] != '#' {
				gj--
				ma[gi][gj] = 'X'
			}
			if gj-1 < 0 {
				break loop
			}
			gd = 'U'
		default:
			panic("invalid direction")
		}
	}

	visited := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if ma[i][j] == 'X' {
				visited++
			}
		}
	}

	fmt.Println(visited)

	return nil
}
