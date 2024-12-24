package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	type ab struct{ a, b string }
	edgeSet := make(map[ab]struct{})
	matrix := make(map[string](map[string]struct{}))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		edge := make([]string, 2)
		_, err := fmt.Sscanf(scanner.Text(), "%2s-%2s", &edge[0], &edge[1])
		if err != nil {
			panic(err)
		}
		slices.Sort(edge)
		a, b := edge[0], edge[1]
		edgeSet[ab{a, b}] = struct{}{}

		if _, ok := matrix[a]; !ok {
			matrix[a] = make(map[string]struct{})
		}
		matrix[a][b] = struct{}{}

		if _, ok := matrix[b]; !ok {
			matrix[b] = make(map[string]struct{})
		}
		matrix[b][a] = struct{}{}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	triplets := 0
	for edge := range edgeSet {
		a, b := edge.a, edge.b
		for c, adjSet := range matrix {
			if a[0] != 't' && b[0] != 't' && c[0] != 't' {
				continue
			}
			if c == a || c == b {
				continue
			}
			_, aPresent := adjSet[a]
			_, bPresent := adjSet[b]
			if aPresent && bPresent {
				triplets++
			}
		}
	}
	triplets /= 3

	fmt.Println(triplets)
}
