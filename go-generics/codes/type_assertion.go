package main


// TYPE ASSERTION OMIT
func NewtonSqrt[T ~float32 | ~float64 ](v T) T {
	var iterations int
	switch (interface{})(v).(type) {
	case float32:
		iterations = 4
	case float64:
		iterations = 5
	default:
		panic(fmt.Sprintf("unexpected type %T", v))
	}
	// Code omitted.
}

type MyFloat float32

var G = NewtonSqrt(MyFloat(64))
// TYPE ASSERTION OMIT