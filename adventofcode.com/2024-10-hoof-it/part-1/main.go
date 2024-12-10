package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
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

	type ij struct{ i, j int }

	topGrid := make([][]map[ij]struct{}, n)
	for i := 0; i < n; i++ {
		topGrid[i] = make([]map[ij]struct{}, m)
	}

	inside := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m
	}

	var topFromIJ func(i, j int) map[ij]struct{}
	topFromIJ = func(i, j int) map[ij]struct{} {
		if topGrid[i][j] != nil {
			return topGrid[i][j]
		}
		top := make(map[ij]struct{})
		if grid[i][j] == 9 {
			top[ij{i, j}] = struct{}{}
		} else {
			neighbors := [...]struct{ i, j int }{
				{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1},
			}
			for _, neighbor := range neighbors {
				if inside(neighbor.i, neighbor.j) && grid[i][j]+1 == grid[neighbor.i][neighbor.j] {
					maps.Insert(top, maps.All(topFromIJ(neighbor.i, neighbor.j)))
				}
			}
		}
		topGrid[i][j] = top
		return top
	}

	answer := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 0 {
				answer += len(topFromIJ(i, j))
			}
		}
	}

	fmt.Println(answer)
}
