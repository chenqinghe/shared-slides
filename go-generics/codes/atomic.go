package main

// OMIT
package atomic 

func AddInt32(addr *int32, delta int32) (new int32)

func AddUint32(addr *uint32, delta uint32) (new uint32)

func AddInt64(addr *int64, delta int64) (new int64)

func AddUint64(addr *uint64, delta uint64) (new uint64)

// OMIT

// INTR OMIT
func Add(addr interface{}, delta interface{}) interface{} {
	switch addr.(type) {
	case nil:
		// some codes
	case *int32:
		// some codes
	case *int64:
		// some codes
	case *uint32:
		// some codes
	case *uint64:
		// some codes
	}

	return nil
}

// INTR OMIT
