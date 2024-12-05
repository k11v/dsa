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
	mulEnabled := true

	scanner := NewScanner(bufio.NewReader(os.Stdin))
	for scanner.Scan() {
		expr := scanner.Token()

		switch expr.Op {
		case "mul":
			if mulEnabled {
				result += expr.Lhs * expr.Rhs
			}
		case "do":
			mulEnabled = true
		case "don't":
			mulEnabled = false
		default:
			panic("unknown operation")
		}
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

	for s.err == nil {
		for s.ch != 'm' && s.ch != 'd' && s.err == nil {
			s.next()
		}

		switch s.ch {
		case 'm':
			s.next()
			if s.scanMul() {
				return true
			}
		case 'd':
			s.next()
			if s.scanDoOrDont() {
				return true
			}
		default:
			continue
		}
	}

	return false
}

func (s *Scanner) scanDoOrDont() bool {
	// 'd' is already scanned.

	if s.ch != 'o' {
		return false
	}
	s.next()

	switch s.ch {
	case '(':
		s.next()
		return s.scanDo()
	case 'n':
		s.next()
		return s.scanDont()
	default:
		return false
	}
}

func (s *Scanner) scanDo() bool {
	// 'd', 'o', '(' are already scanned.

	if s.ch != ')' {
		return false
	}
	s.next()

	op := "do"
	lhs := 0
	rhs := 0

	s.token = ExpressionToken{Op: op, Lhs: lhs, Rhs: rhs}
	return true
}

func (s *Scanner) scanDont() bool {
	// 'd', 'o', 'n' are already scanned.

	if s.ch != '\'' {
		return false
	}
	s.next()

	if s.ch != 't' {
		return false
	}
	s.next()

	if s.ch != '(' {
		return false
	}
	s.next()

	if s.ch != ')' {
		return false
	}
	s.next()

	op := "don't"
	lhs := 0
	rhs := 0

	s.token = ExpressionToken{Op: op, Lhs: lhs, Rhs: rhs}
	return true
}

func (s *Scanner) scanMul() bool {
	// 'm' is already scanned.

	if s.ch != 'u' {
		return false
	}
	s.next()

	if s.ch != 'l' {
		return false
	}
	s.next()

	if s.ch != '(' {
		return false
	}
	s.next()

	lhsRunes := make([]rune, 0)
	for unicode.IsDigit(s.ch) {
		lhsRunes = append(lhsRunes, s.ch)
		s.next()
	}
	if len(lhsRunes) == 0 {
		return false
	}

	if s.ch != ',' {
		return false
	}
	s.next()

	rhsRunes := make([]rune, 0)
	for unicode.IsDigit(s.ch) {
		rhsRunes = append(rhsRunes, s.ch)
		s.next()
	}
	if len(rhsRunes) == 0 {
		return false
	}

	if s.ch != ')' {
		return false
	}
	s.next()

	op := "mul"
	lhs, err := strconv.Atoi(string(lhsRunes))
	if err != nil {
		s.err = err
		return false
	}
	rhs, err := strconv.Atoi(string(rhsRunes))
	if err != nil {
		s.err = err
		return false
	}

	s.token = ExpressionToken{Op: op, Lhs: lhs, Rhs: rhs}
	return true
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
