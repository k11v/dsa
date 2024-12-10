package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	grid := make([][]int, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0, len(line))
		for _, ch := range line {
			h, err := strconv.Atoi(string([]rune{ch}))
			if err != nil {
				panic(err)
			}
			row = append(row, h)
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	n, m := len(grid), len(grid[0])

	scoreGrid := make([][]int, n)
	for i := 0; i < n; i++ {
		scoreGrid[i] = slices.Repeat([]int{-1}, m)
	}

	inside := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m
	}

	var scoreFromIJ func(i, j int) int
	scoreFromIJ = func(i, j int) int {
		if scoreGrid[i][j] != -1 {
			return scoreGrid[i][j]
		}
		score := 0
		if grid[i][j] == 9 {
			score = 1
		} else {
			neighbors := [...]struct{ i, j int }{
				{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1},
			}
			for _, neighbor := range neighbors {
				if inside(neighbor.i, neighbor.j) && grid[i][j]+1 == grid[neighbor.i][neighbor.j] {
					score += scoreFromIJ(neighbor.i, neighbor.j)
				}
			}
		}
		scoreGrid[i][j] = score
		return score
	}

	answer := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				answer += scoreFromIJ(i, j)
			}
		}
	}

	fmt.Println(answer)
}
