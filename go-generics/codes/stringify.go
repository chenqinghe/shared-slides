package main

//  STRINGER OMIT
// Stringer is a type constraint that requires the type argument to have
// a String method and permits the generic function to call String.
// The String method should return a string representation of the value.
type Stringer interface {
	String() string
}

// STRINGER OMIT
