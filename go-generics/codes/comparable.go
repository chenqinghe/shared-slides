package main

// COMPARABLE EXAMPLE OMIT
// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// COMPARABLE EXAMPLE OMIT


// ComparableHasher OMIT
// ComparableHasher is a type constraint that matches all
// comparable types with a Hash method.
type ComparableHasher interface {
	comparable
	Hash() uintptr
}

// ComparableHasher OMIT