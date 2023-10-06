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
	"strconv"
	"strings"
)

// }}}
// {{{ globals

var EXAMPLES = [][]any{
	{"pointers_refs", "pointers and references", func() { pr.Main() }},
	{"structs_ints", "structures and interfaces", si.Main},
}

// }}}
// ---- functions ----
// {{{ func usage()

func usage() {
	fmt.Println("goex - golang examples")
	fmt.Println("")
	fmt.Println("Usage:")
	fmt.Println("")
	fmt.Println("  $ go run main.go [example-id-or-number]")
	fmt.Println("")
	fmt.Println("Examples:")
	fmt.Println("")
	fmt.Println("  $ go run main.go " + EXAMPLES[0][0].(string))
	fmt.Println("  $ go run main.go 1")
	fmt.Println("")
	fmt.Println("Available examples:")
	fmt.Println("")
	top := -1
	for _, v := range EXAMPLES {
		top = max(top, len(v[0].(string)))
	}
	for i, v := range EXAMPLES {
		fmt.Printf("%3d: %s %s %s\n", i+1, v[0], strings.Repeat(".", top-len(v[0].(string))+3), v[1])
	}
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
	exN, err := strconv.Atoi(ex)
	if err == nil && exN >= 1 && exN <= len(EXAMPLES) {
		ex = EXAMPLES[exN-1][0].(string)
	}
	for _, v := range EXAMPLES {
		if v[0] == ex {
			v[2].(func())()
			os.Exit(0)
		}
	}
	fmt.Println("error: unknown example '" + ex + "'")
	os.Exit(1)
}

// }}}
