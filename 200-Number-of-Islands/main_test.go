package main

import (
	"os"
)

// LeetCode

type UnionFind struct{
	p []int
	s []int
}

func NewUnionFind(n int) UnionFind {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return UnionFind{p: p, s: s}
}

func (uf UnionFind) Union(u, v int) {
	u = uf.Find(u)
	v = uf.Find(v)
	if u == v {
		return
	}
	if uf.s[u] < uf.s[v] {
		u, v = v, u
	}
	uf.p[v] = u
	uf.s[u] += uf.s[v]
}

func (uf UnionFind) Find(u int) int {
	if uf.p[u] == u {
		return u
	}
	p := uf.Find(uf.p[u])
	uf.p[u] = p
	return p
}

func numIslands(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	uf := NewUnionFind(n*m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if oi := i-1; oi >= 0 && grid[i][j] == grid[oi][j] {
				uf.Union(i*m+j, oi*m+j)
			}
			if oi := i+1; oi < n && grid[i][j] == grid[oi][j] {
				uf.Union(i*m+j, oi*m+j)
			}
			if oj := j-1; oj >= 0 && grid[i][j] == grid[i][oj] {
				uf.Union(i*m+j, i*m+oj)
			}
			if oj := j+1; oj < m && grid[i][j] == grid[i][oj] {
				uf.Union(i*m+j, i*m+oj)
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if uf.Find(i*m+j) == i*m+j && grid[i][j] == '1' {
				count++
			}
		}
	}

	return count
}
