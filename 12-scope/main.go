package main

import (
	"go/types"
	"os"
)

func main() {
	prog := NewProgram(map[string]string{
		"hello.go": `
			package main
			import "fmt"
			const Pi = 3.14
			func main() {
				for i := 2; i <= 8; i++ {
					fmt.Printf("%d", i)
				}
			}
		`,
		"fmt": `
			package fmt
			func Printf(format string, a ...interface{}) (n int, err error) {
				return
			}
		`,
	})

	pkg, _, _ := prog.LoadPackage("hello.go")
	pkg.Scope().WriteTo(os.Stdout, 0, true)
	pkg.Scope().Parent().WriteTo(os.Stdout, 0, true)
	types.Universe.WriteTo(os.Stdout, 0, true)	// Same as the former one
}