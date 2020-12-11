package ast

import "fmt"

func Print(n Node) {
	print(n, "")
}

func print(n Node, indent string) {
	switch v := n.(type) {
	case *BinaryExpr:
		fmt.Println(indent + "BinaryExpr: " + v.OpKind.String())

		indent += "    "

		print(v.L, indent)
		print(v.R, indent)
	case *LiteralExpr:
		fmt.Println(indent + "LiteralExpr: " + v.Val)
	}
}
