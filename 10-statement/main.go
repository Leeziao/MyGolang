package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	ast.Print(nil, f.Decls[0].(*ast.FuncDecl).Body)
}

var src = `package foo
func main() {
	{}
	{42}
	a := 10
	var b int = 20
	return nil
}
`