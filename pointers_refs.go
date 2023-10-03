//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/nimex
// Date: Tue Oct  3 04:46:23 PM UTC 2023
//
// Compile/run:
//   $ go run pointers_refs.go
//

// {{{ package

package main

// }}}
// {{{ imports

import "fmt"

// }}}
// ---- funcs ----
// {{{ func valChange()

func valChange(x int) {
    prefix := "valChange():"
    fmt.Println(prefix,"x =",x,"/ &x =",&x,"(before)")
    x = 2
    fmt.Println(prefix,"x =",x,"/ &x =",&x,"(after)")
}

// }}}
// {{{ func refChange()

func refChange(x *int) {
    prefix := "refChange():"
    fmt.Println(prefix,"*x =",*x,"/ x =",x,"(before)")
    *x = 2
    fmt.Println(prefix,"*x =",*x,"/ x =",x,"(after)")
}

// }}}
// ---- main ----
// {{{ func main()

func main() {
    var x int = 0
    prefix := "     main():"
    fmt.Println(prefix,"x =",x,"/ &x =",&x)
    valChange(x)
	fmt.Println(prefix,"x =",x,"/ &x =",&x)
    refChange(&x)
	fmt.Println(prefix,"x =",x,"/ &x =",&x)
}

// }}}
