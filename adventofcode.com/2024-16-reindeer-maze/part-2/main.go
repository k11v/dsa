package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type IJ struct{ I, J int }

type IJK struct{ I, J, K int }

type IJKD struct{ I, J, K, D int }

type IJKDP struct {
	I, J, K, D int
	P          IJK
}

type IJKDPHeap []IJKDP

func (h IJKDPHeap) Len() int           { return len(h) }
func (h IJKDPHeap) Less(i, j int) bool { return h[i].D < h[j].D }
func (h IJKDPHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IJKDPHeap) Push(x any)        { *h = append(*h, x.(IJKDP)) }
func (h *IJKDPHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	g := make([][]rune, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		g = append(g, []rune(scanner.Text()))
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	n, m := len(g), len(g[0])

	si, sj := -1, -1
	ei, ej := -1, -1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch g[i][j] {
			case 'S':
				si, sj = i, j
			case 'E':
				ei, ej = i, j
			}
		}
	}
	if si == -1 || sj == -1 || ei == -1 || ej == -1 {
		panic("start or end not found")
	}

	const (
		up = iota
		right
		down
		left
	)
	d := make([][][4]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([][4]int, m)
		for j := 0; j < m; j++ {
			for k := 0; k < 4; k++ {
				d[i][j][k] = math.MaxInt
			}
		}
	}
	p := make([][][4][]IJK, n)
	for i := 0; i < n; i++ {
		p[i] = make([][4][]IJK, m)
	}

	h := &IJKDPHeap{}
	heap.Push(h, IJKDP{si, sj, right, 0, IJK{si, sj, right}})
	for h.Len() > 0 {
		cur := heap.Pop(h).(IJKDP)
		if cur.D > d[cur.I][cur.J][cur.K] {
			continue
		}
		p[cur.I][cur.J][cur.K] = append(p[cur.I][cur.J][cur.K], cur.P)
		if cur.D == d[cur.I][cur.J][cur.K] {
			continue
		}
		d[cur.I][cur.J][cur.K] = cur.D

		neis := []IJKD{
			{cur.I, cur.J, (cur.K + 1) % 4, 1000},
			{cur.I, cur.J, (cur.K - 1 + 4) % 4, 1000},
		}
		switch cur.K {
		case up:
			neis = append(neis, IJKD{cur.I - 1, cur.J, cur.K, 1})
		case right:
			neis = append(neis, IJKD{cur.I, cur.J + 1, cur.K, 1})
		case down:
			neis = append(neis, IJKD{cur.I + 1, cur.J, cur.K, 1})
		case left:
			neis = append(neis, IJKD{cur.I, cur.J - 1, cur.K, 1})
		default:
			panic("unknown direction")
		}
		for _, nei := range neis {
			if g[nei.I][nei.J] != '#' {
				if d[nei.I][nei.J][nei.K] == math.MaxInt {
					heap.Push(h, IJKDP{nei.I, nei.J, nei.K, cur.D + nei.D, IJK{cur.I, cur.J, cur.K}})
				} else if d[nei.I][nei.J][nei.K] == cur.D+nei.D {
					p[nei.I][nei.J][nei.K] = append(p[nei.I][nei.J][nei.K], IJK{cur.I, cur.J, cur.K})
				}
			}
		}
	}

	minD := math.MaxInt
	for k := 0; k < 4; k++ {
		minD = min(minD, d[ei][ej][k])
	}

	visited := make([][][4]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = make([][4]bool, m)
	}
	q := make([]IJK, 0)
	for k := 0; k < 4; k++ {
		if d[ei][ej][k] == minD {
			q = append(q, IJK{ei, ej, k})
		}
	}
	for len(q) > 0 {
		var cur IJK
		cur, q = q[0], q[1:]
		if visited[cur.I][cur.J][cur.K] {
			continue
		}
		visited[cur.I][cur.J][cur.K] = true
		q = append(q, p[cur.I][cur.J][cur.K]...)
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := 0; k < 4; k++ {
				if visited[i][j][k] {
					count++
					break
				}
			}
		}
	}

	fmt.Println(count)
}
