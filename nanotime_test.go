package nanotime

import (
	"testing"
	"time"
)

func BenchmarkTimeNowUnixNano(b *testing.B) {
	for i := 0; i < b.N; i++ {
		time.Now().UnixNano()
	}
}

func BenchmarkNanotime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Now()
	}
}
