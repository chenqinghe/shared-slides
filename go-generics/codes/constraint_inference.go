package main 


// DOUBLE DEFINATION OMIT
// Double returns a new slice that contains all the elements of s, doubled.
func Double[E constraints.Integer](s []E) []E {
	r := make([]E, len(s))
	for i, v := range s {
		r[i] = v + v
	}
	return r
}

var s1 = Double[int]([]int{1,2,3}) // s1 is an int slice with value {2,4,6}

// DOUBLE DEFINATION OMIT


// CALL DOUBLE ERROR OMIT
// MySlice is a slice of ints.
type MySlice []int

// The type of V1 will be []int, not MySlice.
// Here we are using function argument type inference,
// but not constraint type inference.
var V1 = Double(MySlice{1})

// CALL DOUBLE ERROR OMIT


// DOUBLE DEFINED OMIT
// DoubleDefined returns a new slice that contains the elements of s,
// doubled, and also has the same type as s.
func DoubleDefined[S ~[]E, E constraints.Integer](s S) S {
	// Note that here we pass S to make, where above we passed []E.
	r := make(S, len(s))
	for i, v := range s {
		r[i] = v + v
	}
	return r
}
// DOUBLE DEFINED OMIT


// CALL DOUBLE OMIT
// The type of V2 will be MySlice.
var V2 = DoubleDefined[MySlice, int](MySlice{1})

// CALL DOUBLE OMIT


// CALL DOUBLE INFERENCE OMIT
var V3 = DoubleDefined(MySlice{1})
// CALL DOUBLE INFERENCE OMIT

// CALL DOUBLE INFERENCE WHOLE OMIT
var V3 = DoubleDefined[MySlice,int](MySlice{1})
// CALL DOUBLE INFERENCE WHOLE OMIT


// POINTER METHOD OMIT
// Setter is a type constraint that requires that the type
// implement a Set method that sets the value from a string.
type Setter interface {
	Set(string)
}

// FromStrings takes a slice of strings and returns a slice of T,
// calling the Set method to set each returned value.
//
// Note that because T is only used for a result parameter,
// function argument type inference does not work when calling
// this function.
func FromStrings[T Setter](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i].Set(v)
	}
	return result
}
// POINTER METHOD OMIT


