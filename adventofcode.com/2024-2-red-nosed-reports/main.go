package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"flag"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	flag.Parse()

	switch flag.Arg(0) {
	case "1":
		return run1()
	case "2":
		return run2()
	default:
		panic("unknown part")
	}
}

func run1() error {
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

func run2() error {
	safeCount := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		nums, err := intFields(scanner.Text())
		if err != nil {
			return err
		}

		if isSafeWithTolerance(nums) {
			safeCount++
		}
	}
	fmt.Println(safeCount)
	return nil
}

func isSafeWithTolerance(r []int) bool {
	return isSafeWithTolerateAndIncreasing(r, true, false) || isSafeWithTolerateAndIncreasing(r, true, true)
}

func isSafeWithTolerateAndIncreasing(r []int, tolerate bool, increasing bool) bool {
	if len(r) <= 1 {
		return true
	}

	for i := 1; i < len(r); i++ {
		directionOK := r[i] - r[i-1] > 0 == increasing

		d := abs(r[i] - r[i-1])
		diffOK := d >= 1 && d <= 3 

		if !directionOK || !diffOK {
			if tolerate {
				withoutPrev := append(append([]int(nil), r[:i-1]...), r[i:]...)
				withoutCurr := append(append([]int(nil), r[:i]...), r[i+1:]...)
				return isSafeWithTolerateAndIncreasing(withoutPrev, false, increasing) || isSafeWithTolerateAndIncreasing(withoutCurr, false, increasing)
			}
			return false
		}
	}

	return true
}

func isSafeWithSkippedIndex(r []int, skippedIndex *int) bool {
	if len(r) <= 1 {
		return true
	}

	increasing := r[1] - r[0] > 0

	var prev int
	for i, curr := range r {
		if skippedIndex != nil && i == *skippedIndex {
			continue
		}

		if i == 0 || i == 1 && skippedIndex != nil && *skippedIndex == 0 {
			prev = curr
			continue
		}

		directionOK := curr - prev > 0 == increasing

		d := abs(curr - prev)
		diffOK := d >= 1 && d <= 3 

		if !directionOK || !diffOK {
			if skippedIndex != nil {
				return false
			}
			prevIndex := i-1
			currIndex := i
			return isSafeWithSkippedIndex(r, &prevIndex) || isSafeWithSkippedIndex(r, &currIndex)
		}
		prev = curr
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

