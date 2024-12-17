package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var ra, rb, rc int
	_, err := fmt.Scanf("Register A: %d\n", &ra)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("Register B: %d\n", &rb)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("Register C: %d\n", &rc)
	if err != nil {
		panic(err)
	}

	var programLine string
	_, err = fmt.Scanf("\nProgram: %s\n", &programLine)
	if err != nil {
		panic(err)
	}
	parts := strings.Split(programLine, ",")
	program := make([]int, len(parts))
	for i, part := range parts {
		program[i], err = strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
	}

	literal := func(operand int) int {
		return operand
	}
	combo := func(operand int) int {
		switch operand {
		case 0, 1, 2, 3:
			return operand
		case 4:
			return ra
		case 5:
			return rb
		case 6:
			return rc
		case 7:
			panic("combo operand 7 is undefined")
		default:
			panic("unknown operand")
		}
	}

	out := make([]int, 0)
	i := 0
	for i < len(program) {
		opcode := program[i]
		i++
		operand := program[i]
		i++

		switch opcode {
		case 0:
			ra = ra / (1 << combo(operand))
		case 1:
			rb = rb ^ literal(operand)
		case 2:
			rb = combo(operand) % 8
		case 3:
			if ra != 0 {
				i = literal(operand)
			}
		case 4:
			rb = rb ^ rc
			_ = operand
		case 5:
			out = append(out, combo(operand)%8)
		case 6:
			rb = ra / (1 << combo(operand))
		case 7:
			rc = ra / (1 << combo(operand))
		default:
			panic("unknown opcode")
		}
	}

	for i, o := range out {
		if i != 0 {
			fmt.Print(",")
		}
		fmt.Print(o)
	}
	fmt.Print("\n")
}
