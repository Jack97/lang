package main

import (
	"github.com/Jack97/lang/ast"
	"github.com/Jack97/lang/parser"
)

func main() {
	var p parser.Parser
	p.Init("2 * (10.234 + 3) * 10")

	tree, err := p.Parse()

	ast.Print(tree)

	if err != nil {
		parser.PrintError(err)
	}
}
