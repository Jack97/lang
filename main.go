package main

import (
	"github.com/Jack97/lang/ast"
	"github.com/Jack97/lang/parser"
)

func main() {
	var p parser.Parser
	p.Init("2 * (10 + 3) * 10 + 5")

	tree := p.Parse()

	ast.Print(tree)
}
