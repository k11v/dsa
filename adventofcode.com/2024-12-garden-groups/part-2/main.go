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

	edge := func(other, cell struct{ i, j int }) bool {
		return !inside(other.i, other.j) || grid[other.i][other.j] != grid[cell.i][cell.j]
	}

	eq := func(a, b struct{ i, j int }) bool {
		return inside(a.i, a.j) && inside(b.i, b.j) && grid[a.i][a.j] == grid[b.i][b.j]
	}

	countCorners := func(i, j int) int {
		type ij struct{ i, j int }
		c := ij{i, j}
		u, d, l, r := ij{i - 1, j}, ij{i + 1, j}, ij{i, j - 1}, ij{i, j + 1}
		ul, dr, dl, ur := ij{i - 1, j - 1}, ij{i + 1, j + 1}, ij{i + 1, j - 1}, ij{i - 1, j + 1}
		return boolToInt(edge(l, c) && edge(u, c)) +
			boolToInt(edge(u, c) && edge(r, c)) +
			boolToInt(edge(r, c) && edge(d, c)) +
			boolToInt(edge(d, c) && edge(l, c)) +
			boolToInt(eq(u, c) && eq(l, c) && !eq(ul, c)) +
			boolToInt(eq(d, c) && eq(r, c) && !eq(dr, c)) +
			boolToInt(eq(d, c) && eq(l, c) && !eq(dl, c)) +
			boolToInt(eq(u, c) && eq(u, c) && !eq(ur, c))
	}

	var dfs func(i, j int) (area, corners int)
	dfs = func(i, j int) (area, corners int) {
		if visited[i][j] {
			return 0, 0
		}
		visited[i][j] = true

		area = 1
		corners = countCorners(i, j)
		fmt.Println(i, j, corners)

		neighbors := []struct{ i, j int }{{i - 1, j}, {i + 1, j}, {i, j - 1}, {i, j + 1}}
		for _, neighbor := range neighbors {
			if inside(neighbor.i, neighbor.j) {
				if grid[neighbor.i][neighbor.j] == grid[i][j] {
					a, c := dfs(neighbor.i, neighbor.j)
					area += a
					corners += c
				}
			}
		}

		return area, corners
	}

	total := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			a, c := dfs(i, j)
			total += a * c
		}
	}

	fmt.Println(total)
	return
}

func boolToInt(v bool) int {
	if v {
		return 1
	} else {
		return 0
	}
}
