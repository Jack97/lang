package token

type Token int

const (
	ILLEGAL Token = iota
	EOF

	INT

	ADD
	SUB
	MUL
	DIV

	LPAREN
	RPAREN
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	INT: "INT",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	DIV: "/",

	LPAREN: "(",
	RPAREN: ")",
}

func (tok Token) String() string {
	return tokens[tok]
}
