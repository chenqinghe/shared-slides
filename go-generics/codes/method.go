package main

import "fmt"

// START OMIT
type P[T any] struct{}

func (p *P[T]) Say[S any](v S) { fmt.Println(v) }

type T struct{}

func (t *T) Say[S any](v S) { fmt.Println(v) }

func main() {
	var t T
	t.Say[int](1)

	/** output:
	method.go:8:19: methods cannot have type parameters
	method.go:8:20: invalid AST: method must have no type parameters
	method.go:12:16: methods cannot have type parameters
	method.go:12:17: invalid AST: method must have no type parameters
	*/
}

// END OMIT
