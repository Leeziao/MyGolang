package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strconv"
)

func main() {
	expr, _ := parser.ParseExpr(`1*2+3 +y`)
	environ := map[string]float64 {
		"x": 100,
	}
	fmt.Println(Eval(expr, environ))
}

func Eval(exp ast.Expr, vars map[string]float64) float64 {
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp, vars)
	case *ast.BasicLit:
		f, _ := strconv.ParseFloat(exp.Value, 64)
		return f
	case *ast.Ident:
		val, ok := vars[exp.Name]
		if ok {
			return val
		}
		panic(fmt.Sprintf("Variable %s not declared", exp.Name))
	}
	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr, vars map[string]float64) float64 {
	switch exp.Op {
	case token.ADD:
		return Eval(exp.X, vars) + Eval(exp.Y, vars)
	case token.MUL:
		return Eval(exp.X, vars) * Eval(exp.Y, vars)
	}
	return 0
}
