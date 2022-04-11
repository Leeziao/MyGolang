package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", src, parser.AllErrors)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Packge(%s)\n", f.Name.Name)
	for _, s := range f.Imports {
		fmt.Printf("import: %s\n", s.Path.Value)
	}
	for _, decl := range f.Decls {
		fmt.Printf("decl: %T\n", decl)
	}
	_ = ast.Bad
}

var src = `package pkgname

import ("a"; "b")
type SomeType int
const PI = 3.14
var Length = 1

func main() {}
`