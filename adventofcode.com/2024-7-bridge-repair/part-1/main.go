package main

import (
	"bufio"
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
	answer := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		resultString, numbersString, _ := strings.Cut(line, ": ")
		stringNumbers := strings.Split(numbersString, " ")

		result, err := strconv.Atoi(resultString)
		if err != nil {
			return err
		}
		numbers := make([]int, 0, len(stringNumbers))
		for _, stringNumber := range stringNumbers {
			number, err := strconv.Atoi(stringNumber)
			if err != nil {
				return err
			}
			numbers = append(numbers, number)
		}

		if check(numbers, result) {
			answer += result
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	fmt.Println(answer)
	return nil
}

func check(numbers []int, result int) bool {
	if result < 0 {
		return false
	}
	if len(numbers) == 0 {
		return result == 0
	}

	if result%numbers[len(numbers)-1] == 0 {
		if check(numbers[:len(numbers)-1], result/numbers[len(numbers)-1]) {
			return true
		}
	}
	return check(numbers[:len(numbers)-1], result-numbers[len(numbers)-1])
}
