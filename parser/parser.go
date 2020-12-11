package parser

import (
	"fmt"

	"github.com/Jack97/lang/ast"
	"github.com/Jack97/lang/scanner"
	"github.com/Jack97/lang/token"
)

type Parser struct {
	scanner scanner.Scanner
	Errors  ErrorList

	tok token.Token
	pos int
	lit string
}

func (p *Parser) Init(src string) {
	eh := func(pos int, msg string) {
		p.Errors.Add(pos, msg)
	}

	p.scanner.Init(src, eh)

	p.next()
}

func (p *Parser) next() {
	p.tok, p.pos, p.lit = p.scanner.Scan()
}

func (p *Parser) error(pos int, msg string) {
	p.Errors.Add(pos, msg)
}

func (p *Parser) Parse() ast.Node {
	expr := p.parseExpr()

	if p.tok != token.EOF {
		p.error(p.pos, fmt.Sprintf("expected '%s', got '%s'", token.EOF, p.tok))
	}

	return expr
}

func (p *Parser) parseExpr() ast.Expr {
	return p.parseTermExpr()
}

func (p *Parser) parseTermExpr() ast.Expr {
	left := p.parseFactorExpr()

	for p.tok == token.ADD || p.tok == token.SUB {
		opKind, opPos := p.tok, p.pos

		p.next()

		right := p.parseFactorExpr()

		left = &ast.BinaryExpr{
			L:      left,
			OpKind: opKind,
			OpPos:  opPos,
			R:      right,
		}
	}

	return left
}

func (p *Parser) parseFactorExpr() ast.Expr {
	left := p.parsePrimaryExpr()

	for p.tok == token.MUL || p.tok == token.DIV {
		opKind, opPos := p.tok, p.pos

		p.next()

		right := p.parsePrimaryExpr()

		left = &ast.BinaryExpr{
			L:      left,
			OpKind: opKind,
			OpPos:  opPos,
			R:      right,
		}
	}

	return left
}

func (p *Parser) parsePrimaryExpr() ast.Expr {
	tok, pos, lit := p.tok, p.pos, p.lit

	p.next()

	if tok == token.INT {
		return &ast.LiteralExpr{
			Kind:   tok,
			ValPos: pos,
			Val:    lit,
		}
	}

	p.error(pos, fmt.Sprintf("expected operand, got '%s'", p.tok))

	return &ast.BadExpr{
		FromPos: pos,
		ToPos:   p.pos,
	}
}
