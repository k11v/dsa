package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type rawItem struct{ used, free int }

type zipItem struct{ id, used int }

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	raw := make([]rawItem, 0)
	r := bufio.NewReader(os.Stdin)
	for {
		ch1, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if ch1 == '\n' {
			break
		}

		ch2, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				ch2 = '0'
			} else {
				return err
			}
		}
		if ch2 == '\n' {
			ch2 = '0'
		}

		raw = append(raw, rawItem{used: digitToInt(ch1), free: digitToInt(ch2)})
	}

	zip := make([]zipItem, 0)
	i, j := 0, len(raw)-1
outer:
	for {
		zip = append(zip, zipItem{id: i, used: raw[i].used})
		if i == j {
			break outer
		}
		for raw[i].free != 0 {
			take := min(raw[i].free, raw[j].used)
			raw[j].used -= take
			raw[i].free -= take
			zip = append(zip, zipItem{id: j, used: take})
			if raw[j].used == 0 {
				j--
				if i == j {
					break outer
				}
			}
		}
		i++
	}

	answer := 0
	k := 0
	for _, zi := range zip {
		for l := 0; l < zi.used; l++ {
			answer += k * zi.id
			k++
		}
	}

	fmt.Println(answer)
	return nil
}

func digitToInt(digit rune) int {
	switch digit {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	default:
		panic("not a digit")
	}
}
