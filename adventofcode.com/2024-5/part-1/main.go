package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)

	rules := make(map[int](map[int]struct{}))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		var l, r int
		_, err := fmt.Sscanf(line, "%d|%d", &l, &r)
		if err != nil {
			return err
		}

		if _, ok := rules[l]; !ok {
			rules[l] = make(map[int]struct{})
		}
		rules[l][r] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	s := 0
	for scanner.Scan() {
		line := scanner.Text()

		stringNumbers := strings.Split(line, ",")
		if len(stringNumbers) == 0 {
			return errors.New("empty update")
		}

		numbers := make([]int, 0, len(stringNumbers))
		for _, stringNumber := range stringNumbers {
			number, err := strconv.Atoi(stringNumber)
			if err != nil {
				return err
			}

			numbers = append(numbers, number)
		}

		if isValid(numbers, rules) {
			s += numbers[len(numbers)/2]
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println(s)

	return nil
}

func isValid(numbers []int, rules map[int](map[int]struct{})) bool {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			ni, nj := numbers[i], numbers[j]
			if _, oppositeRuleExists := rules[nj][ni]; oppositeRuleExists {
				return false
			}
		}
	}
	return true
}
