package ast

import "fmt"

func Print(n Node) {
	print(n, "")
}

func print(n Node, indent string) {
	switch x := n.(type) {
	case *BinaryExpr:
		fmt.Println(indent + "BinaryExpr [" + x.OpKind.String() + "]:")

		indent += "    "

		print(x.L, indent)
		print(x.R, indent)
	case *LiteralExpr:
		fmt.Println(indent + "LiteralExpr [" + x.Val + "]:")
	case *ParenExpr:
		fmt.Println(indent + "ParenExpr:")

		indent += "    "

		print(x.Expr, indent)
	}
}
