package scanner

import (
	"fmt"

	"github.com/Jack97/lang/token"
)

type ErrorHandler func(pos int, msg string)

type Scanner struct {
	src     string
	eh      ErrorHandler
	pos     int
	nextPos int
	char    rune
}

func (s *Scanner) Init(src string, eh ErrorHandler) {
	s.src = src
	s.eh = eh
	s.next()
}

func (s *Scanner) next() {
	s.pos = s.nextPos

	if s.pos >= len(s.src) {
		s.char = -1
	} else {
		s.char = rune(s.src[s.pos])
		s.nextPos++
	}
}

func (s *Scanner) error(pos int, msg string) {
	if s.eh != nil {
		s.eh(pos, msg)
	}
}

func (s *Scanner) skipWhitespace() {
	for s.char == ' ' || s.char == '\t' || s.char == '\n' || s.char == '\r' {
		s.next()
	}
}

func (s *Scanner) Scan() (tok token.Token, pos int, lit string) {
	s.skipWhitespace()

	pos = s.pos

	if isDecimal(s.char) {
		s.next()

		for isDecimal(s.char) {
			s.next()
		}

		tok, lit = token.INT, s.src[pos:s.pos]
	} else {
		switch s.char {
		case -1:
			tok = token.EOF
		case '+':
			tok = token.ADD
		case '-':
			tok = token.SUB
		case '*':
			tok = token.MUL
		case '/':
			tok = token.DIV
		case '(':
			tok = token.LPAREN
		case ')':
			tok = token.RPAREN
		default:
			tok, lit = token.ILLEGAL, string(s.char)
			s.error(pos, fmt.Sprintf("illegal character %#U", s.char))
		}

		s.next()
	}

	return
}

func isDecimal(char rune) bool {
	return '0' <= char && char <= '9'
}
