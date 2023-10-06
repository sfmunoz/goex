//
// vim: set foldmethod=marker:
//
// URL:  https://github.com/sfmunoz/goex
// Date: Tue Oct  3 04:46:23 PM UTC 2023
//

// {{{ package

package pointers_refs

// }}}
// {{{ imports

import "fmt"

// }}}
// ---- funcs (private) ----
// {{{ func valChange()

func valChange(x int) {
	prefix := "valChange():"
	fmt.Println(prefix, "x =", x, "/ &x =", &x, "(before)")
	x = 2
	fmt.Println(prefix, "x =", x, "/ &x =", &x, "(after)")
}

// }}}
// {{{ func refChange()

func refChange(x *int) {
	prefix := "refChange():"
	fmt.Println(prefix, "*x =", *x, "/ x =", x, "(before)")
	*x = 2
	fmt.Println(prefix, "*x =", *x, "/ x =", x, "(after)")
}

// }}}
// ---- funcs (public) ----
// {{{ func Main()

func Main() {
	var x int = 0
	prefix := "     main():"
	fmt.Println(prefix, "x =", x, "/ &x =", &x)
	valChange(x)
	fmt.Println(prefix, "x =", x, "/ &x =", &x)
	refChange(&x)
	fmt.Println(prefix, "x =", x, "/ &x =", &x)
}

// }}}
