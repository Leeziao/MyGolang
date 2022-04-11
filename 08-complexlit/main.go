package main

import (
	"go/ast"
	"go/parser"
)
var exprs_str = []string{
	"[...]int{1, 2, 3}",
	"func() {}",
	"func(a int, b string) (result bool)",
}

func main() {
	for _, expr_str := range exprs_str{
		expr, _ := parser.ParseExpr(expr_str)
		ast.Print(nil, expr)
	}
}