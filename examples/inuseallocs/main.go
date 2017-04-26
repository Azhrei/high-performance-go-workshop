// A simple example to demonstrate the difference between alloc_count
// and inuse_count.
package main

import (
	"math/rand"
	"runtime"
	"github.com/pkg/profile"	// from dave.cheney.net
)

const count = 100000

var y []byte

func main() {
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()
	y = allocate()
	runtime.GC()		// Force GC execution
}

// allocate returns one row, a []byte, from 'count' number of rows.
func allocate() []byte {
	var x [][]byte
	for i := 0; i < count; i++ {
		x = append(x, makeByteSlice())
	}
	return x[0]
}

// makeByteSlice returns a []byte of a size up to 16k.
func makeByteSlice() []byte {
	return make([]byte, rand.Intn(2^14))	// Up to 16k
}
