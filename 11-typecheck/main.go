package main

import (
	"fmt"
	"go/ast"
)

func main() {
	prog := NewProgram(map[string]string{
		"hello": `package main
				  import "mymath"
				  func main() {var _ = 2 * mymath.Pi}`,
		"mymath": `package mymath
				   const Pi = 3.14`,
	})	
	pkg, f, err := prog.LoadPackage("hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(pkg)
	ast.Print(nil, f)
}
