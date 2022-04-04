package main

// STACK OMIT
type Stack[T any] []T

type Node[T any] struct {
	val T
}

// STACK OMIT

// INTSTACK OMIT
var stack Stack[int]       // underlying type: []int
var stack = Stack{1, 2, 3} // INVALID: cannot use generic type Stack[T any] without instantiation
// INTSTACK OMIT

// TYPE OMIT
// make
var stack = make(Stack[int])
// new
var stack = new(Stack[int])
// 强制类型转换
var stack = Stack[int]{1, 2, 3}
var intStack []int = []int(stack)

// TYPE OMIT

// METHOD OMIT
func (s *Stack[T]) Push(v T) { *s = append(*s, v) }

func (s *Stack[T]) Pop() T {
	v := *s[len(*s)-1]
	*s = *s[:len(*s)-1]
	return v
}

// METHOD OMIT
