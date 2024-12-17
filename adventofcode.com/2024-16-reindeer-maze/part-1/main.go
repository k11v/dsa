package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type IJKD struct{ I, J, K, D int }

type IJKDHeap []IJKD

func (h IJKDHeap) Len() int           { return len(h) }
func (h IJKDHeap) Less(i, j int) bool { return h[i].D < h[j].D }
func (h IJKDHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IJKDHeap) Push(x any)        { *h = append(*h, x.(IJKD)) }
func (h *IJKDHeap) Pop() any {
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

	h := &IJKDHeap{}
	heap.Push(h, IJKD{si, sj, right, 0})
	for h.Len() > 0 {
		cur := heap.Pop(h).(IJKD)
		if cur.D >= d[cur.I][cur.J][cur.K] {
			continue
		}
		d[cur.I][cur.J][cur.K] = cur.D
		if cur.I == ei && cur.J == ej {
			break
		}

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
					heap.Push(h, IJKD{nei.I, nei.J, nei.K, cur.D + nei.D})
				}
			}
		}
	}

	minD := math.MaxInt
	for k := 0; k < 4; k++ {
		minD = min(minD, d[ei][ej][k])
	}
	fmt.Println(minD)
}
