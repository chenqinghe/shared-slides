package main

import "fmt"

/**
// PRINT OMIT
func Print(s []T) { // T is a type parameter
	for _, v := range s {
		fmt.Println(v)
	}
}

// PRINT OMIT
**/

// GEN OMIT
func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// GEN OMIT

func main() {
	// CALL OMIT
	Print[int]([]int{1, 2, 3, 4})
	Print[int32]([]int32{1, 2, 3, 4})
	// CALL OMIT
}

/**
// CALL MULTI PARAMS OMIT
func Index[K comparable,V any](m map[K]V, k K) V {
	return m[k]
}

Index[int,string](map[int]string{1:"hello"}, 1)
// CALL MULTI PARAMS OMIT


*/

func dummy() {
	// CALL TYPE OMIT
	Print([]int{1, 2, 3})
	// CALL TYPE OMIT
	// COMPLETE CALL OMIT
	Print[int]([]int{1, 2, 3})
	// COMPLETE CALL OMIT
}
