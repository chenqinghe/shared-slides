package main

// ANY OMIT
// any is a type alias of empty interface
type any = interface{}

// ANY OMIT

// PRINT OMIT
// Print prints the elements of any slice.
// Print has a type parameter T and has a single (non-type)
// parameter s which is a slice of that type parameter.
func Print[T interface{}](s []T) {
	// same as above
}

// PRINT OMIT
