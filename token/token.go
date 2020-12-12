package token

type Token int

const (
	ILLEGAL Token = iota
	EOF

	INT
	FLOAT

	PLUS
	MINUS
	STAR
	SLASH

	LPAREN
	RPAREN
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",

	INT:   "INT",
	FLOAT: "FLOAT",

	PLUS:  "+",
	MINUS: "-",
	STAR:  "*",
	SLASH: "/",

	LPAREN: "(",
	RPAREN: ")",
}

func (tok Token) String() string {
	return tokens[tok]
}
