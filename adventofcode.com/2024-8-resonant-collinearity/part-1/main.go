package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type ij struct{ i, j int }

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

// \r\n is not handled.
func run() error {
	n, m := 0, 0
	antennas := make(map[rune][]ij)

	r := bufio.NewReader(os.Stdin)
	i, j := 0, 0
	for {
		ch, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				if j != 0 {
					m = max(m, j)
					i++
					j = 0
				}
				break
			}
			return err
		}
		if ch == '\n' {
			m = max(m, j)
			i++
			j = 0
			continue
		}
		if ch != '.' {
			antennas[ch] = append(antennas[ch], ij{i, j})
		}
		j++
	}
	n = i

	antinodes := make(map[ij]struct{})
	for _, antennaLocs := range antennas {
		for p := 0; p < len(antennaLocs); p++ {
			for q := p + 1; q < len(antennaLocs); q++ {
				antennaLoc1 := antennaLocs[p]
				antennaLoc2 := antennaLocs[q]

				antinodeLoc1 := ij{
					2*antennaLoc1.i - antennaLoc2.i,
					2*antennaLoc1.j - antennaLoc2.j,
				}
				if antinodeLoc1.i >= 0 && antinodeLoc1.i < n && antinodeLoc1.j >= 0 && antinodeLoc1.j < m {
					antinodes[antinodeLoc1] = struct{}{}
				}

				antinodeLoc2 := ij{
					2*antennaLoc2.i - antennaLoc1.i,
					2*antennaLoc2.j - antennaLoc1.j,
				}
				if antinodeLoc2.i >= 0 && antinodeLoc2.i < n && antinodeLoc2.j >= 0 && antinodeLoc2.j < m {
					antinodes[antinodeLoc2] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(antinodes))
	return nil
}
