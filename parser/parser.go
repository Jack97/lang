package parser

import (
	"fmt"

	"github.com/Jack97/lang/ast"
	"github.com/Jack97/lang/scanner"
	"github.com/Jack97/lang/token"
)

type Parser struct {
	scanner scanner.Scanner
	errors  ErrorList
	tok     token.Token
	pos     int
	lit     string
}

func (p *Parser) Init(src string) {
	eh := func(pos int, msg string) {
		p.errors.Add(pos, msg)
	}

	p.scanner.Init(src, eh)

	p.next()
}

func (p *Parser) next() {
	p.tok, p.pos, p.lit = p.scanner.Scan()
}

func (p *Parser) error(pos int, msg string) {
	p.errors.Add(pos, msg)
}

func (p *Parser) Parse() (ast.Node, error) {
	expr := p.parseExpr()

	if p.tok != token.EOF {
		p.error(p.pos, fmt.Sprintf("expected '%s', got '%s'", token.EOF, p.tok))
	}

	if len(p.errors) > 0 {
		return expr, p.errors
	}

	return expr, nil
}

func (p *Parser) parseExpr() ast.Expr {
	return p.parseBinaryExpr(0)
}

func (p *Parser) parseBinaryExpr(prevPrec int) ast.Expr {
	left := p.parsePrimaryExpr()

	for {
		opKind, opPos := p.tok, p.pos

		prec := 0

		switch opKind {
		case token.STAR, token.SLASH:
			prec = 2
		case token.PLUS, token.MINUS:
			prec = 1
		}

		if prec == 0 || prec <= prevPrec {
			return left
		}

		p.next()

		right := p.parseBinaryExpr(prec)

		left = &ast.BinaryExpr{
			L:      left,
			OpKind: opKind,
			OpPos:  opPos,
			R:      right,
		}
	}
}

func (p *Parser) parsePrimaryExpr() ast.Expr {
	tok, pos, lit := p.tok, p.pos, p.lit

	p.next()

	if tok == token.LPAREN {
		expr := p.parseExpr()
		rparenPos := p.pos

		if p.tok != token.RPAREN {
			p.error(rparenPos, fmt.Sprintf("expected '%s', got '%s'", token.RPAREN, p.tok))
		}

		p.next()

		return &ast.ParenExpr{
			LparenPos: pos,
			Expr:      expr,
			RparenPos: rparenPos,
		}
	}

	if tok == token.INT || tok == token.FLOAT {
		return &ast.LiteralExpr{
			Kind:   tok,
			ValPos: pos,
			Val:    lit,
		}
	}

	p.error(pos, fmt.Sprintf("expected operand, got '%s'", tok))

	return &ast.BadExpr{
		FromPos: pos,
		ToPos:   p.pos,
	}
}
