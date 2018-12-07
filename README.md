# High Performance Time Formatter

This package is built for highly optimized applications and packages
like loggers. The API allows you to preallocate the required memory,
this in turn leads to zero memory allocations within the formatter.

`timef` is up to two times faster compared to `time.Time.Format`.


## Benchmark

```
goos: darwin
goarch: amd64
pkg: github.com/mrcrgl/timef
Benchmark_Format-4          	10000000	       166 ns/op	      64 B/op	       2 allocs/op
Benchmark_FormatBytes-4     	10000000	       131 ns/op	      32 B/op	       1 allocs/op
Benchmark_FormatRFC3339-4   	10000000	       126 ns/op	      32 B/op	       1 allocs/op
Benchmark_WriteRFC3339-4    	20000000	       100 ns/op	       0 B/op	       0 allocs/op
Benchmark_TimeFormat-4      	10000000	       225 ns/op	      32 B/op	       1 allocs/op
```