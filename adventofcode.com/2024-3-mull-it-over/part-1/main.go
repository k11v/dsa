package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"unicode"
)

func main() {
	result := 0

	scanner := NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		expr := scanner.Token()
		if expr.Op != "mul" {
			panic("operation is not mul")
		}
		result += expr.Lhs * expr.Rhs
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(result)
}

const end = -1

type ExpressionToken struct {
	Op  string
	Lhs int
	Rhs int
}

type Scanner struct {
	r *bufio.Reader

	ch          rune
	err         error
	initialized bool
	token       ExpressionToken
}

func NewScanner(r *bufio.Reader) *Scanner {
	return &Scanner{r: r}
}

func (s *Scanner) Scan() bool {
	if !s.initialized {
		s.next()
		s.initialized = true
	}

	var token *ExpressionToken

	for token == nil && s.err == nil {
		for s.ch != 'm' && s.ch != end {
			s.next()
		}
		s.next()

		if s.ch != 'u' {
			continue
		}
		s.next()

		if s.ch != 'l' {
			continue
		}
		s.next()

		if s.ch != '(' {
			continue
		}
		s.next()

		lhsRunes := make([]rune, 0)
		for unicode.IsDigit(s.ch) {
			lhsRunes = append(lhsRunes, s.ch)
			s.next()
		}
		if len(lhsRunes) == 0 {
			continue
		}

		if s.ch != ',' {
			continue
		}
		s.next()

		rhsRunes := make([]rune, 0)
		for unicode.IsDigit(s.ch) {
			rhsRunes = append(rhsRunes, s.ch)
			s.next()
		}
		if len(rhsRunes) == 0 {
			continue
		}

		if s.ch != ')' {
			continue
		}
		s.next()

		op := "mul"
		lhs, err := strconv.Atoi(string(lhsRunes))
		if err != nil {
			s.err = err
			continue
		}
		rhs, err := strconv.Atoi(string(rhsRunes))
		if err != nil {
			s.err = err
			continue
		}

		token = &ExpressionToken{Op: op, Lhs: lhs, Rhs: rhs}
	}

	if token != nil {
		s.token = *token
		return true
	}
	return false
}

func (s *Scanner) Token() ExpressionToken {
	return s.token
}

func (s *Scanner) Err() error {
	if errors.Is(s.err, io.EOF) {
		return nil
	}
	return s.err
}

func (s *Scanner) next() {
	r, _, err := s.r.ReadRune()
	if err != nil {
		s.err = err
		s.ch = end
		return
	}
	s.ch = r
}
