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
	cx, cy = sx, sy
	for cx != ex || cy != ey {
		rx, ry := cx, cy
		for cx != ex || cy != ey {
			neighbors := [...]xy{
				{cx, cy + 1},
				{cx + 1, cy},
				{cx, cy - 1},
				{cx - 1, cy},
			}
			for _, n := range neighbors {
				if inside(n.x, n.y) && g[n.y][n.x] == '.' && d[cy][cx]+1 == d[n.y][n.x] {
					cx, cy = n.x, n.y
					break
				}
			}
			cheatActiveDistance := abs(cx-rx) + abs(cy-ry)
			if cheatActiveDistance > 20 {
				continue
			}
			cheatDistance := fairDistance - (d[cy][cx] - d[ry][rx] - cheatActiveDistance)
			if fairDistance-cheatDistance >= 100 {
				goodCheatCount++
			}
		}
		cx, cy = rx, ry
		neighbors := [...]xy{
			{cx, cy + 1},
			{cx + 1, cy},
			{cx, cy - 1},
			{cx - 1, cy},
		}
		for _, n := range neighbors {
			if inside(n.x, n.y) && g[n.y][n.x] == '.' && d[cy][cx]+1 == d[n.y][n.x] {
				cx, cy = n.x, n.y
				break
			}
		}
	}

	fmt.Println(goodCheatCount)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
