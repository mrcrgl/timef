# High Performance Time Formatter

This package is built for highly optimized applications and packages
like loggers. The API allows you to preallocate the required memory,
this in turn leads to zero memory allocations within the formatter.

`timef` is up to two times faster compared to `time.Time.Format`.
