package main 

import (
	"fmt"
	"strconv"
)


type Settable int

func (p *Settable) Set(s string) {
	i, _ := strconv.Atoi(s) // real code should not ignore the error
	*p = Settable(i)
}

// SETTER2 OMIT
type Setter2[B any] interface {
	Set(string)
	*B // non-interface type constraint element
}

func FromStrings2[T any, PT Setter2[T]](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		// The type of &result[i] is *T which is in the type set
		// of Setter2, so we can convert it to PT.
		p := PT(&result[i]) 
		p.Set(v) // PT has a Set method.
	}
	return result
}
// SETTER2 OMIT


func main(){
// FromStrings2 takes two type parameters.
	// The second parameter must be a pointer to the first.
	// Settable is as above.
// SETTER2 CALL OMIT
	nums := FromStrings2[Settable, *Settable]([]string{"1", "2"})
	// Now nums is []Settable{1, 2}.
	fmt.Println(nums)
// SETTER2 CALL OMIT


/**
// SETTER3 CALL OMIT
nums := FromStrings2[Settable]([]string{"1", "2"})
// SETTER3 CALL OMIT
*/

}