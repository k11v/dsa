package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

type (
	IJ  struct{ I, J int }
	IJD struct{ I, J, D int }
)

type IJDHeap []IJD

func (h IJDHeap) Len() int           { return len(h) }
func (h IJDHeap) Less(i, j int) bool { return h[i].D < h[j].D }
func (h IJDHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IJDHeap) Push(x any)        { *h = append(*h, x.(IJD)) }
func (h *IJDHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	si, sj := 0, 0
	ei, ej := 70, 70
	n, m := 71, 71
	g := make([][]bool, n)
	for i := range len(g) {
		g[i] = make([]bool, m)
		for j := range len(g[i]) {
			g[i][j] = true
		}
	}

	fallen := 0
	scanner := bufio.NewScanner(os.Stdin)
	for fallen < 1024 {
		if !scanner.Scan() {
			panic("eof")
		}
		var i, j int
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d", &i, &j)
		if err != nil {
			panic(err)
		}
		g[i][j] = false
		fallen++
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	d := make([][]int, n)
	for i := range len(d) {
		d[i] = make([]int, m)
		for j := range len(d[i]) {
			d[i][j] = math.MaxInt
		}
	}

	h := &IJDHeap{}
	heap.Push(h, IJD{si, sj, 0})
	for h.Len() > 0 {
		cur := heap.Pop(h).(IJD)
		if cur.D >= d[cur.I][cur.J] {
			continue
		}
		d[cur.I][cur.J] = cur.D
		if cur.I == ei && cur.J == ej {
			break
		}

		inside := func(i, j int) bool {
			return i >= 0 && i < n && j >= 0 && j < m
		}

		neis := []IJD{
			{cur.I - 1, cur.J, 1},
			{cur.I, cur.J + 1, 1},
			{cur.I + 1, cur.J, 1},
			{cur.I, cur.J - 1, 1},
		}
		for _, nei := range neis {
			if inside(nei.I, nei.J) && g[nei.I][nei.J] {
				if d[nei.I][nei.J] == math.MaxInt {
					heap.Push(h, IJD{nei.I, nei.J, cur.D + nei.D})
				}
			}
		}
	}

	fmt.Println(d[ei][ej])
}
