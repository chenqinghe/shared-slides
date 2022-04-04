package main


// ZERO VALUE OMIT
type Optional[T any] struct { p *T }
func (o Optional[T]) Val() T {
	if o.p != nil {
		return *o.p
	}
	// how do we return default value of type T ?
}

// ZERO VALUE OMIT