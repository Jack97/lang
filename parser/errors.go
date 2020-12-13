package parser

import (
	"fmt"
)

type Error struct {
	Pos int
	Msg string
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Pos, e.Msg)
}

type ErrorList []*Error

func (e *ErrorList) Add(pos int, msg string) {
	*e = append(*e, &Error{Pos: pos, Msg: msg})
}

func (e ErrorList) Error() string {
	switch len(e) {
	case 0:
		return "no errors"
	case 1:
		return e[0].Error()
	case 2:
		return fmt.Sprintf("%s (and 1 more error)", e[0])
	}

	return fmt.Sprintf("%s (and %d more errors)", e[0], len(e)-1)
}

func PrintError(err error) {
	if list, ok := err.(ErrorList); ok {
		for _, e := range list {
			fmt.Printf("%s\n", e)
		}
	} else if err != nil {
		fmt.Printf("%s\n", err)
	}
}
