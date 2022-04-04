package main



// INTERFACE OMIT

type ReadWriteCloser interface{
	ReadWriter 
	Close() error
}

// INTERFACE OMIT