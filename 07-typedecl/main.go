package main

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// Type    Expr   (*Ident, *ParenExpr, *SelectorExpr, *StarExpr, or any of the *XxxTypes)
func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		panic("Error")
	}
	for _, decl := range f.Decls {
		if decl_gen, ok := decl.(*ast.GenDecl); ok {
			ast.Print(nil, decl_gen.Specs[0])
		}
	}

}

var src = `package foo
type Int1 int
type Int2 pkg.int
type IntPtr *int
type Mystruct struct {a, b int}
type IntStringMap map[int]String
`