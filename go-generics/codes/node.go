package main

import "fmt"

// NODE OMIT
type Node[T any] struct {
	next *Node[T]
	val  T
}

func main() {
	var a = Node[string]{next: &Node[string]{}, val: "ssss"}
	fmt.Println(a)
}

// NODE OMIT
