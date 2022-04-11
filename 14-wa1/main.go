package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/constant"
	"go/types"
	"os"

	"golang.org/x/tools/go/ssa"
)

func runFunc(fn *ssa.Function) {
	fmt.Println("--- runFunc begin ---")
	defer fmt.Println("--- runFunc end ---")

	if len(fn.Blocks) > 0 {
		for blk := fn.Blocks[0]; blk != nil; {
			blk = runFuncBlock(fn, blk)
		}
	}
}

func runFuncBlock(fn *ssa.Function, block *ssa.BasicBlock) (nextBlock *ssa.BasicBlock) {
	for _, ins := range block.Instrs {
		switch ins := ins.(type) {
		case *ssa.Call:
			doCall(ins)
		// case *ssa.Return:
		// 	doReturn(ins)
		// default:
		// 	doUnknown(ins)
		}
	}
	return nil
}

func doCall(ins *ssa.Call) {
	switch {
		case ins.Call.Method == nil :
			switch callFn := ins.Call.Value.(type) {
			case *ssa.Builtin:
				callBuiltin(callFn, ins.Call.Args...)
		}
	}
}

func callBuiltin(fn *ssa.Builtin, args ...ssa.Value) {
	switch fn.Name() {
	case "println":
		var buf bytes.Buffer
		for i := 0; i < len(args); i++ {
			if i > 0 {
				buf.WriteRune(' ')
			}
			switch arg := args[i].(type) {
			case *ssa.Const:
				if t, ok := arg.Type().(*types.Basic); ok {
					switch t.Kind() {
					case types.String:
						fmt.Fprintf(&buf, "%s", constant.StringVal(arg.Value))
					default:
					}
				}
			}
		}
		buf.WriteRune('\n')
		os.Stdout.Write(buf.Bytes())
	default:
	}
}

func main() {
	prog := NewProgram(map[string]string{
		"test": src,
	})

	pkg, f, info, err := prog.LoadPackage("test")
	if err != nil {
		panic(err)
	}

	var ssaProg = ssa.NewProgram(prog.fset, ssa.SanityCheckFunctions)
	var ssaPkg = ssaProg.CreatePackage(pkg, []*ast.File{f}, info, true)
	
	ssaPkg.Build()
	ssaPkg.Func("main").WriteTo(os.Stdout)

	runFunc(ssaPkg.Func("main"))
}

var src = `
package main

func main() {
	println("Hello World")
}
`