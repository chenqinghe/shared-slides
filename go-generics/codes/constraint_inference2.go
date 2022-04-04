package main

import (
	"strconv"
)

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

// POINTER METHOD CALL OMIT
// Settable is an integer type that can be set from a string.
type Settable int

// Set sets the value of *p from a string.
func (p *Settable) Set(s string) {
	i, _ := strconv.Atoi(s) // real code should not ignore the error
	*p = Settable(i)
}

func main() {
	// INVALID
	nums := FromStrings[Settable]([]string{"1", "2"})
	_ = nums
	// Here we want nums to be []Settable{1, 2}.
	// ...
}
// POINTER METHOD CALL OMIT