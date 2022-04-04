package main

// STACK OMIT
type Stack []interface{}

func (s *Stack) Push(v interface{}) { *s = append(*s, v) }

func (s *Stack) Pop() interface{} {
	v := *s[len(*s)-1]
	*s = *s[:len(*s)-1]
	return v
}

// STACK OMIT
