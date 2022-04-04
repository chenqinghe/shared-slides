package main

// MIN OMIT
func Min[T any](s []T) T {
	r := s[0] // assume that s is not empty
	for _, v := range s {
		if v < r { // INVALID
			r = v
		}
	}
	return r
}

// MIN OMIT



// OPERATOR CONSTRAINT OMIT

// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// OPERATOR CONSTRAINT OMIT

// OPERATOR MIN OMIT
func Min[T Ordered](s []T) T {
	r := s[0] // assume that s is not empty
	for _, v := range s {
		if v < r { // INVALID
			r = v
		}
	}
	return r
}

// OPERATOR MIN OMIT


// OPERATOR MIN VALID OMIT
func Min[T Ordered](s []T) T {
	r := s[0] // assume that s is not empty
	for _, v := range s {
		if v < r { // VALID!
			r = v
		}
	}
	return r
}

// OPERATOR MIN VALID OMIT