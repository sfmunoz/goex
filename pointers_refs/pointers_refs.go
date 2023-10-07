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
// ---- types ----
// {{{ type Data struct

type Data struct {
	val int
}

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
// {{{ func(d Data) valChange()

func (d Data) valChange(val int) {
	d.val = val
}

// }}}
// {{{ func(d *Data) refChange()

func (d *Data) refChange(val int) {
	d.val = val // alt: (*d).val = val
}

// }}}
// {{{ func(d Data) String()

func (d Data) String() string {
	return fmt.Sprintf("val = %d", d.val)
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
	fmt.Println("--------")
	d := Data{1}
	fmt.Printf("%s %s -- initial\n", prefix, d)
	d.valChange(2)
	fmt.Printf("%s %s -- d.valChange(2) applied\n", prefix, d)
	d.refChange(3)
	fmt.Printf("%s %s -- d.refChange(3) applied\n", prefix, d)
}

// }}}
