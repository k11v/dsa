package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var raIn, rbIn, rcIn int
	_, err := fmt.Scanf("Register A: %d\n", &raIn)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("Register B: %d\n", &rbIn)
	if err != nil {
		panic(err)
	}
	_, err = fmt.Scanf("Register C: %d\n", &rcIn)
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

	for raIn = 0; ; raIn++ {
		if raIn%1000000 == 0 {
			fmt.Println(raIn)
		}
		programOut := make([]int, 0)
		ra, rb, rc := raIn, rbIn, rcIn

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

		for i := 0; i < len(program); {
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
				programOut = append(programOut, combo(operand)%8)
			case 6:
				rb = ra / (1 << combo(operand))
			case 7:
				rc = ra / (1 << combo(operand))
			default:
				panic("unknown opcode")
			}
		}

		same := true
		if len(program) != len(programOut) {
			same = false
		} else {
			for i := 0; i < len(program); i++ {
				if program[i] != programOut[i] {
					same = false
					break
				}
			}
		}
		if same {
			break
		}
	}

	fmt.Println(raIn)
}
