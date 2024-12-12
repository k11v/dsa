package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	grid := make([][]rune, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		row := []rune(line)
		grid = append(grid, row)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	n, m := len(grid), len(grid[0])

	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	inside := func(i, j int) bool {
		return i >= 0 && i < n && j >= 0 && j < m
	}

	var dfs func(i, j int) (area, perimeter int)
	dfs = func(i, j int) (area, perimeter int) {
		if visited[i][j] {
			return 0, 0
		}
		visited[i][j] = true

		area = 1
		perimeter = 0

		neighbors := []struct{ i, j int }{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
		for _, neighbor := range neighbors {
			if inside(neighbor.i, neighbor.j) {
				if grid[neighbor.i][neighbor.j] == grid[i][j] {
					a, p := dfs(neighbor.i, neighbor.j)
					area += a
					perimeter += p
				} else {
					perimeter += 1
				}
			} else {
				perimeter += 1
			}
		}

		return area, perimeter
	}

	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a, p := dfs(i, j)
			total += a * p
		}
	}

	fmt.Println(total)
	return
}
