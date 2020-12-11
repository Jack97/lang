package ast

import "github.com/Jack97/lang/token"

type Node interface {
	StartPos() int
	EndPos() int
}

type Expr interface {
	Node
	exprNode()
}

type LiteralExpr struct {
	Kind   token.Token
	ValPos int
	Val    string
}

type BinaryExpr struct {
	L      Expr
	OpKind token.Token
	OpPos  int
	R      Expr
}

type ParenExpr struct {
	LparenPos int
	Expr      Expr
	RparenPos int
}

type BadExpr struct {
	FromPos int
	ToPos   int
}

func (n *LiteralExpr) StartPos() int { return n.ValPos }
func (n *BinaryExpr) StartPos() int  { return n.L.StartPos() }
func (n *ParenExpr) StartPos() int   { return n.LparenPos }
func (n *BadExpr) StartPos() int     { return n.FromPos }

func (n *LiteralExpr) EndPos() int { return n.ValPos + len(n.Val) }
func (n *BinaryExpr) EndPos() int  { return n.R.EndPos() }
func (n *ParenExpr) EndPos() int   { return n.RparenPos + 1 }
func (n *BadExpr) EndPos() int     { return n.ToPos }

func (*LiteralExpr) exprNode() {}
func (*BinaryExpr) exprNode()  {}
func (*ParenExpr) exprNode()   {}
func (*BadExpr) exprNode()     {}
