# GoEx: Golang Examples

This repository holds [Go](https://go.dev/) snippets created while I'm recalling and learning the language

## References

* [Effective Go](https://go.dev/doc/effective_go)
* [A Tour of Go](https://go.dev/tour/)
* [Go Programming](https://www.youtube.com/watch?v=CF9S4QZuV30) ([Derek Banas](https://www.youtube.com/@derekbanas))
  * [Cheat sheet](https://www.newthinktank.com/2015/02/go-programming-tutorial/)

## Packages vs Modules

From [How to Write Go Code](https://go.dev/doc/code):

* **Package:**
  * Go programs are organized into packages.
  * A package is a collection of source files in the same directory that are compiled together.
  * Functions, types, variables, and constants defined in one source file are visible to all other source files within the same package.
* **Module:**
  * A repository contains one or more modules.
  * A module is a collection of related Go packages that are released together.
  * A Go repository typically contains only one module, located at the root of the repository.
  * A file named go.mod there declares the module path: the import path prefix for all packages within the module

## Core

- [main.go](https://github.com/sfmunoz/goex/blob/main/main.go): executable, command-line parsing, [type assertions (any / interface{})](https://go.dev/tour/methods/15) ...
- [make_vs_new.go](https://github.com/sfmunoz/goex/blob/main/make_vs_new/make_vs_new.go): make vs new
  - [Why would I make() or new()?](https://stackoverflow.com/questions/9320862/why-would-i-make-or-new)
  - [Golang New vs Make](https://medium.com/learn-code/golang-new-vs-make-8a4dbd84e92b)
- [pointers_refs.go](https://github.com/sfmunoz/goex/blob/main/pointers_refs/pointers_refs.go): pointers and references
  - [Should I define methods on values or pointers?](https://go.dev/doc/faq#methods_on_values_or_pointers):
    - Does the method need to modify the receiver?
    - It will be much cheaper to use a pointer receiver.
    - If some of the methods of the type must have pointer receivers, the rest should too, so the method set is consistent regardless of how the type is used.
  - [Why do T and *T have different method sets?](https://go.dev/doc/faq#different_method_sets)
    - The method set of a type T consists of all methods with receiver type T
    - That of the corresponding pointer type *T consists of all methods with receiver *T or T.
    - That means the method set of *T includes that of T, but not the reverse.
- [structs_ints.go](https://github.com/sfmunoz/goex/blob/main/structs_ints/structs_ints.go): structs and interfaces
