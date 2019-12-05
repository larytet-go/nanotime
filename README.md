# nanotime

This is a wrapper around not exported runtime.nanotime(). The API is 2x faster than time.Now()

    BenchmarkTimeNowUnixNano-4   	30000000	        56.4 ns/op
    BenchmarkNanotime-4          	50000000	        23.8 ns/op

