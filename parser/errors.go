package parser

type Error struct {
	Pos int
	Msg string
}

type ErrorList []*Error

func (e *ErrorList) Add(pos int, msg string) {
	*e = append(*e, &Error{Pos: pos, Msg: msg})
}

func (e *ErrorList) Reset() {
	*e = (*e)[0:0]
}
