package main

// TYPE OMIT
type Integer interface {
	int
}

// TYPE OMIT

// INVALID OMIT
// EmbeddedParameter is INVALID.
type EmbeddedParameter[T any] interface {
	T // INVALID: may not list a plain type parameter
}
// INVALID OMIT


// APPRO OMIT
type MyInt int 

type AnyInteger interface {
	~int 
}
// APPRO OMIT


// ALL INVALID OMIT

type MyString string 

// ApproximateMyString is INVALID.
type ApproximateMyString interface {
	~MyString // INVALID: underlying type of MyString is not MyString
}

// ApproximateParameter is INVALID.
type ApproximateParameter[T any] interface {
	~T // INVALID: T is a type parameter
}

type MyInterface {}

// ApproximateInterface is INVALID.
type ApproximateInterface interface {
	~MyInterface // INVALID: T is an interface type
}

// ALL INVALID OMIT




// ONION EXAMPLE OMIT

// PredeclaredSignedInteger is a constraint that matches the
// five predeclared signed integer types.
type PredeclaredSignedInteger interface {
	int | int8 | int16 | int32 | int64
}

// SignedInteger is a constraint that matches any signed integer type.
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// ONION EXAMPLE OMIT


// StringableSignedInteger OMIT
type StringableSignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	String() string
}
// StringableSignedInteger OMIT


// COMPOSITE TYPE OMIT
type byteseq interface{
	string | []byte
}

func Join[T byteseq](a []T, sep T) (ret T) {
	// some checks...
	
	n := len(sep) * (len(a) - 1)
	for _, v := range a {
		n += len(v)
	}

	b := make([]byte, n)

	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}

	return T(b)
}
// COMPOSITE TYPE OMIT


// COMPOSITE STRUCT FIELD OMIT
// structField is a type constraint whose type set consists of some
// struct types that all have a field named x.
type structField interface {
	struct { a int; x int } |
		struct { b int; x float64 } |
		struct { c int; x uint64 }
}

// This function is INVALID.
func IncrementX[T structField](p *T) {
	v := p.x // INVALID: type of p.x is not the same for all types in set
	v++
	p.x = v
}
// COMPOSITE STRUCT FIELD OMIT



// TYPE CONVERSION OMIT
type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Convert[To, From integer](from From) To {
	to := To(from)
	if From(to) != from {
		panic("conversion out of range")
	}
	return to
}

// TYPE CONVERSION OMIT


// UNTYPED CONSTANTS OMIT
type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func Add10[T integer](s []T) {
	for i, v := range s {
		s[i] = v + 10 // OK: 10 can convert to any integer type
	}
}

// This function is INVALID.
func Add1024[T integer](s []T) {
	for i, v := range s {
		s[i] = v + 1024 // INVALID: 1024 not permitted by int8/uint8
	}
}
// UNTYPED CONSTANTS OMIT


// EMBED INTERSECTION OMIT
// Addable is types that support the + operator.
type Addable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~complex64 | ~complex128 |
		~string
}

// Byteseq is a byte sequence: either string or []byte.
type Byteseq interface {
	~string | ~[]byte
}

// AddableByteseq is a byte sequence that supports +.
// This is every type that is both Addable and Byteseq.
// In other words, just the type set ~string.
type AddableByteseq interface {
	Addable
	Byteseq
}
// EMBED INTERSECTION OMIT



// EMBED UNION OMIT
// Signed is a constraint with a type set of all signed integer
// types.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned is a constraint with a type set of all unsigned integer
// types.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

// Integer is a constraint with a type set of all integer types.
type Integer interface {
	Signed | Unsigned
}
// EMBED UNION OMIT

// INTERFACE IN UNION OMIT
type Stringish interface {
	string | fmt.Stringer
}

func ToString[T Stringish](v T) string {
	switch x := (interface{})(v).(type) {
	case string:
	  return x
  
	case fmt.Stringer:
	  return x.String()
	}
  
	panic("impossible")
  }
// INTERFACE IN UNION OMIT

// EMPTY TYPE SETS OMIT
// Unsatisfiable is an unsatisfiable constraint with an empty type set.
// No predeclared types have any methods.
// If this used ~int | ~float32 the type set would not be empty.
type Unsatisfiable interface {
	int | float32
	String() string
}
// EMPTY TYPE SETS OMIT