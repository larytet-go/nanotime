package nanotime

import (
	_ "unsafe"
)

//go:noescape
//go:linkname nanotime runtime.nanotime
func nanotime() int64

// Now returns a timestamp
func Now() int64 {
	return nanotime()
}
