package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	TypeDecl()
}

func TypeDecl() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src_type, parser.AllErrors)
	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}
	for _, decl := range f.Decls {
		if v, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range v.Specs {
				fmt.Printf("%T\n", spec)
				ast.Print(nil, spec)
			}
		}
	}
}
var src_type = `package foo
import "pkg1"
import _pkg2 "pkg2"

type MyInt1 int
type MyInt2 = int

const PI = 3.14
const E float64 = 2.718

var Pi = 3.14
`