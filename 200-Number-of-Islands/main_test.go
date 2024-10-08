package main

// LeetCode

type IJ struct { i, j int }

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	count := 0

	n, m := len(grid), len(grid[0])
	outsideQueue := make([]IJ, 0)
	outsideQueue = append(outsideQueue, IJ{0, 0})
	insideQueue := make([]IJ, 0)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	for len(outsideQueue) != 0 {
		var r IJ
		r, outsideQueue = outsideQueue[0], outsideQueue[1:]
		if visited[r.i][r.j] {
			continue
		}

		insideQueue = append(insideQueue, r)
		for len(insideQueue) != 0 {
			var p IJ
			p, insideQueue = insideQueue[0], insideQueue[1:]
			if visited[p.i][p.j] {
				continue
			}
			visited[p.i][p.j] = true

			neighbors := []IJ{
				{p.i-1, p.j},
				{p.i+1, p.j},
				{p.i, p.j-1},
				{p.i, p.j+1},
			}
			for _, q := range neighbors {
				if !(q.i >= 0 && q.i < n && q.j >= 0 && q.j < m) {
					continue
				}
				if visited[q.i][q.j] {
					continue
				}

				if grid[q.i][q.j] == grid[p.i][p.j] {
					insideQueue = append(insideQueue, q)
				} else {
					outsideQueue = append(outsideQueue, q)
				}
			}
		}

		if grid[r.i][r.j] == '1' {
			count++
		}
	}

	return count
}
