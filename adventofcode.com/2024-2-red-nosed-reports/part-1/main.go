package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	safeCount := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums, err := intFields(scanner.Text())
		if err != nil {
			return err
		}

		if isSafe(nums) {
			safeCount++
		}
	}
	fmt.Println(safeCount)
	return nil
}

func isSafe(r []int) bool {
	if len(r) <= 1 {
		return true
	}

	increasing := r[1] - r[0] > 0

	for i := 1; i < len(r); i++ {
		directionOK := r[i] - r[i-1] > 0 == increasing

		d := abs(r[i] - r[i-1])
		diffOK := d >= 1 && d <= 3 

		if !directionOK || !diffOK {
			return false
		}
	}

	return true
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}

func intFields(s string) ([]int, error) {
	strs := strings.Fields(s)
	nums := make([]int, 0, len(strs))
	for _, str := range strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}
