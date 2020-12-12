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

func (s *Scanner) peek() rune {
	pos := s.pos + 1

	if pos >= len(s.src) {
		return -1
	}

	return rune(s.src[pos])
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

func (s *Scanner) scanDigits() {
	for isDigit(s.char) {
		s.next()
	}
}

func (s *Scanner) scanNumber() (token.Token, string) {
	pos := s.pos
	tok := token.ILLEGAL

	// integer part
	if s.char != '.' {
		tok = token.INT
		s.scanDigits()
	}

	// fractional part
	if s.char == '.' && isDigit(s.peek()) {
		tok = token.FLOAT
		s.next()
		s.scanDigits()
	}

	lit := s.src[pos:s.pos]

	return tok, lit
}

func (s *Scanner) Scan() (tok token.Token, pos int, lit string) {
	s.skipWhitespace()

	pos = s.pos

	switch char := s.char; {
	case isDigit(char) || char == '.' && isDigit(s.peek()):
		tok, lit = s.scanNumber()
	default:
		s.next() // make progress

		switch char {
		case -1:
			tok = token.EOF
		case '+':
			tok = token.PLUS
		case '-':
			tok = token.MINUS
		case '*':
			tok = token.STAR
		case '/':
			tok = token.SLASH
		case '(':
			tok = token.LPAREN
		case ')':
			tok = token.RPAREN
		default:
			tok, lit = token.ILLEGAL, string(s.char)
			s.error(pos, fmt.Sprintf("illegal character %#U", s.char))
		}
	}

	return
}

func isDigit(char rune) bool {
	return '0' <= char && char <= '9'
}
