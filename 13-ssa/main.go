package main

import (
	"fmt"
	"go/ast"
	"go/types"
	"os"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/interp"
)

func main() {
	prog := NewProgram(map[string]string{
		"main": main_src,
		"runtime": runtime_src,
	})
	prog.LoadPackage("main")
	prog.LoadPackage("runtime")

	var ssaProg = ssa.NewProgram(prog.fset, ssa.SanityCheckFunctions)
	var ssaMainPkg *ssa.Package

	for name, pkg := range prog.pkgs {
		var ssaPkg = ssaProg.CreatePackage(pkg, []*ast.File{prog.ast[name]}, prog.infos[name], true)
		if name == "main" {
			ssaMainPkg = ssaPkg
		}
	}
	ssaProg.Build()
	ssaProg.Package(prog.pkgs["main"]).Func("main").WriteTo(os.Stdout)

	exitCode := interp.Interpret(ssaMainPkg, 0, &types.StdSizes{8, 8}, "main", []string{})
	if exitCode != 0 {
		fmt.Println("exitCode: ", exitCode)
	}

	// ssaPkg.WriteTo(os.Stdout)
	// ssaPkg.Func("init").WriteTo(os.Stdout)
	// ssaPkg.Func("main").WriteTo(os.Stdout)
}

var main_src = `
package main

var s = "hello ssa"

func main() {
	for i := 0; i < 3; i++ {
		println(s)
	}
}
`

var runtime_src = `
package runtime

type errorString string
func (e errorString) RuntimeError() {}
func (e errorString) Error() string { return "runtime error: " + string(e)}

type Error interface {
	error
	RuntimeError()
}

`