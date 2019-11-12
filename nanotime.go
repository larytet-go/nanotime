package nanotime

import (
	_ "unsafe" // I need for go:linkname
)

// Now returns a timestamp
func Now() int64 {
	return nanotime()
}

//go:noescape
//go:linkname nanotime runtime.nanotime
func nanotime() int64
