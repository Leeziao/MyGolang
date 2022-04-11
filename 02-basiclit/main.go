package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	var lit1234 = &ast.BasicLit{
		Kind: token.INT,
		Value: "1234",
	}
	ast.Print(nil, lit1234)

	expr1, _ := parser.ParseExpr(`2345`)
	ast.Print(nil, expr1)

	ast.Print(nil, ast.NewIdent("x"))
	expr2, _ := parser.ParseExpr(`x`)
	ast.Print(nil, expr2)
}