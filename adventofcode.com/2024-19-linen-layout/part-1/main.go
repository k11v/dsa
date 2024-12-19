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

	var possible func(design string) bool
	possible = func(design string) bool {
		if design == "" {
			return true
		}
		for _, towel := range towels {
			subdesign, found := strings.CutPrefix(design, towel)
			if !found {
				continue
			}
			if possible(subdesign) {
				return true
			}
		}
		return false
	}

	count := 0
	for scanner.Scan() {
		design := scanner.Text()
		if possible(design) {
			count++
		}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(count)
}
