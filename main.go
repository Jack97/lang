package main

import (
	"fmt"

	"github.com/Jack97/lang/ast"
	"github.com/Jack97/lang/parser"
)

func main() {
	var p parser.Parser
	p.Init("2 * 10 + 3 * 10")

	tree := p.Parse()

	ast.Print(tree)

	for _, e := range p.Errors {
		fmt.Println(e.Pos, e.Msg)
	}
}
