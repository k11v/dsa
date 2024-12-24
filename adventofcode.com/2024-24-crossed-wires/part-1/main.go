package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	values := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		var wire string
		var bit int
		_, err := fmt.Sscanf(line, "%3s: %d", &wire, &bit)
		if err != nil {
			panic(err)
		}

		values[wire] = bit
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	type expression struct{ left, op, right string }
	expressions := make(map[string]expression)
	for scanner.Scan() {
		line := scanner.Text()
		var left, op, right, result string
		_, err := fmt.Sscanf(line, "%s %s %s -> %s", &left, &op, &right, &result)
		if err != nil {
			panic(err)
		}
		expressions[result] = expression{left, op, right}
	}
	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	var get func(wire string) int
	get = func(wire string) int {
		if value, ok := values[wire]; ok {
			return value
		}
		expression := expressions[wire]
		left := get(expression.left)
		op := expression.op
		right := get(expression.right)

		var value int
		switch op {
		case "AND":
			value = left & right
		case "OR":
			value = left | right
		case "XOR":
			value = left ^ right
		default:
			panic("unknown operation")
		}

		values[wire] = value
		return value
	}

	wiresWithZ := make([]string, 0)
	for wire := range expressions {
		if strings.HasPrefix(wire, "z") {
			wiresWithZ = append(wiresWithZ, wire)
		}
	}
	slices.Sort(wiresWithZ)

	num := 0
	for i, wire := range wiresWithZ {
		num += get(wire) * (1 << i)
	}

	fmt.Println(num)
}
