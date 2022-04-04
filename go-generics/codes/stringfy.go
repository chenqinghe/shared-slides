package main


func main(){
	Stringify([]int{1,2,3})
}

// START OMIT
// Stringify convert a slice of type T to a string slice.
func Stringify[T any](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}

// END OMIT