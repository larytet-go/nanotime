package nanotime

//go:noescape
//go:linkname nanotime runtime.nanotime
func nanotime() int64

func Now() int64 {
	return nanotime()
}
