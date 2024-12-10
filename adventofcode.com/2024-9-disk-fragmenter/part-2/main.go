package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type item struct{ id, used, free int }

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	disk := make([][]item, 0)
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

		disk = append(disk, []item{{id: len(disk), used: digitToInt(ch1), free: digitToInt(ch2)}})
	}

	for j := len(disk) - 1; j >= 1; j-- {
		js := 0
		need := disk[j][js].used
		for i := 0; i < j; i++ {
			ie := len(disk[i]) - 1
			if disk[i][ie].free >= need {
				disk[i] = append(disk[i], disk[j][js])
				disk[i][ie+1].free = disk[i][ie].free - need
				disk[i][ie].free = 0

				disk[j][js].id = -1
				disk[j][js].used = 0
				disk[j][js].free = disk[j][js].free + need

				break
			}
		}
	}

	answer := 0
	k := 0
	for _, items := range disk {
		for _, itm := range items {
			for l := 0; l < itm.used; l++ {
				answer += k * itm.id
				k++
			}
			k += itm.free
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
