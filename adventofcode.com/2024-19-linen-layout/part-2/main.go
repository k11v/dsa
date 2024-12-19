package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var towels []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if scanner.Text() == "" {
			break
		}
		towels = append(towels, strings.Split(line, ", ")...)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	memo := make(map[string]int)
	var possible func(design string) int
	possible = func(design string) int {
		if p, ok := memo[design]; ok {
			return p
		}
		if design == "" {
			memo[design] = 1
			return 1
		}
		total := 0
		for _, towel := range towels {
			subdesign, found := strings.CutPrefix(design, towel)
			if !found {
				continue
			}
			total += possible(subdesign)
		}
		memo[design] = total
		return total
	}

	count := 0
	for scanner.Scan() {
		design := scanner.Text()
		count += possible(design)
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(count)
}
