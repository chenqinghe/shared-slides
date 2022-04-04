
// MAP OMIT
func Map[F, T any](s []F, f func(F) T) []T { ... }
// MAP OMIT


// MAP CALL OMIT
var s []int
f := func(i int) int64 { return int64(i) }
var r []int64
// Specify both type arguments explicitly.
r = Map[int, int64](s, f)
// Specify just the first type argument, for F,
// and let T be inferred.
r = Map[int](s, f)
// Don't specify any type arguments, and let both be inferred.
r = Map(s, f)

// MAP CALL OMIT


// ERR OMIT

func Double[T interface{ ~int }](v T) T { 
	return v * 2 
}

Double(1) // 

// ERR OMIT


// CANNOT INFERENCE OMIT

func NewValue[T any]()T {
	var t T 
	return t 
}

println(NewValue()) // cannot infer T
// CANNOT INFERENCE OMIT


// MAP2 OMIT
// Map calls the function f on every element of the slice s,
// returning a new slice of the results.
func Map[F, T any](s []F, f func(F) T) []T {
	r := make([]T, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

strs := Map([]int{1, 2, 3}, strconv.Itoa)
// MAP2 OMIT

// PAIR OMIT
// NewPair returns a pair of values of the same type.
func NewPair[F any](f1, f2 F) *Pair[F] { ... }
// PAIR OMIT