package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	g := make([][]rune, 0)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		g = append(g, []rune(scanner.Text()))
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}
	w, h := len(g[0]), len(g)

	sx, sy := -1, -1
	ex, ey := -1, -1
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			switch g[y][x] {
			case 'S':
				sx, sy = x, y
				g[y][x] = '.'
			case 'E':
				ex, ey = x, y
				g[y][x] = '.'
			}
		}
	}
	if sx == -1 || sy == -1 || ex == -1 || ey == -1 {
		panic("start or end not found")
	}

	d := make([][]int, h)
	for y := 0; y < h; y++ {
		d[y] = make([]int, w)
		for x := 0; x < w; x++ {
			d[y][x] = -1
		}
	}
	d[sy][sx] = 0

	inside := func(x, y int) bool {
		return x >= 0 && x < w && y >= 0 && y < h
	}
	type xy struct{ x, y int }
	cx, cy := sx, sy
	for cx != ex || cy != ey {
		neighbors := [...]xy{
			{cx, cy + 1},
			{cx + 1, cy},
			{cx, cy - 1},
			{cx - 1, cy},
		}
		for _, n := range neighbors {
			if inside(n.x, n.y) && g[n.y][n.x] == '.' && d[n.y][n.x] == -1 {
				d[n.y][n.x] = d[cy][cx] + 1
				cx, cy = n.x, n.y
				break
			}
		}
	}

	goodCheatCount := 0
	fairDistance := d[ey][ex]
	type xyxy struct{ x1, y1, x2, y2 int }
	cx, cy = sx, sy
	for cx != ex || cy != ey {
		neighbors := [...]xyxy{
			{cx, cy + 1, cx, cy + 2},
			{cx + 1, cy, cx + 2, cy},
			{cx, cy - 1, cx, cy - 2},
			{cx - 1, cy, cx - 2, cy},
		}
		for _, n := range neighbors {
			if inside(n.x1, n.y1) && inside(n.x2, n.y2) && g[n.y1][n.x1] == '#' && g[n.y2][n.x2] == '.' {
				cheatDistance := fairDistance - (d[n.y2][n.x2] - d[cy][cx] - 2)
				if fairDistance-cheatDistance >= 100 {
					goodCheatCount++
				}
			}
		}
		for _, n := range neighbors {
			if inside(n.x1, n.y1) && g[n.y1][n.x1] == '.' && d[cy][cx]+1 == d[n.y1][n.x1] {
				cx, cy = n.x1, n.y1
				break
			}
		}
	}

	fmt.Println(goodCheatCount)
}
