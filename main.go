//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/goex
// Date: Fri Oct  6 03:31:34 PM UTC 2023
//
// Compile/run:
//   $ go run main.go
//

// {{{ package

package main

// }}}
// {{{ imports

import (
	"flag"
	"fmt"
	pr "github.com/sfmunoz/goex/pointers_refs"
	si "github.com/sfmunoz/goex/structs_ints"
	"os"
)

// }}}
// ---- functions ----
// {{{ func usage()

func usage() {
	fmt.Println("goex - golang examples")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("  $ go run main.go [example]")
	fmt.Println("")
	fmt.Println("Available examples:")
	fmt.Println("")
	fmt.Println("  pointers_refs: pointers and references")
	fmt.Println("   structs_ints: structures and interfaces")
	fmt.Println("")
	fmt.Println("Reference:")
	fmt.Println("")
	fmt.Println("  https://github.com/sfmunoz/goex/")
	fmt.Println("")
}

// }}}
// ---- main ----
// {{{ func main()

func main() {
	flag.Parse()
	tot := flag.NArg()
	if tot < 1 {
		usage()
		os.Exit(0)
	}
	if tot > 1 {
		fmt.Println("error: only one example can be specified")
		os.Exit(1)
	}
	ex := flag.Arg(0)
	switch ex {
	case "pointers_refs":
		pr.Main()
		os.Exit(0)
	case "structs_ints":
		si.Main()
		os.Exit(0)
	}
	fmt.Println("error: unknown example '" + ex + "'")
	os.Exit(1)
}

// }}}
