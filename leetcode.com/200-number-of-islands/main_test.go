package main

// LeetCode

func numIslands(g [][]byte) int {
	if len(g) == 0 || len(g[0]) == 0 {
		return 0
	}

	count := 0

	n, m := len(g), len(g[0])
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i >= 0 && i < n && j >= 0 && j < m && g[i][j] == '1' && !visited[i][j] {
			visited[i][j] = true
			dfs(i-1, j)
			dfs(i+1, j)
			dfs(i, j-1)
			dfs(i, j+1)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if g[i][j] == '1' && !visited[i][j] {
				dfs(i, j)
				count++
			}
		}
	}

	return count
}
