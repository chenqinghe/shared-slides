package main

// STR OMIT
// Stringify calls the String method on each element of s,
// and returns the results.
func Stringify[T Stringer](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

// STR OMIT

// PRT OMIT
// Print2 has two type parameters and two non-type parameters.
func Print2[T1, T2 any](s1 []T1, s2 []T2) { /**...*/ }

// PRT OMIT

// CONCAT OMIT
type Stringer interface{ String() string }

type Plusser interface{ Plus(string) string }

func ConcatTo[S Stringer, P Plusser](s []S, p []P) []string {
	r := make([]string, len(s))
	for i, v := range s {
		r[i] = p[i].Plus(v.String())
	}
	return r
}

// CONCAT OMIT
