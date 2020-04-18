## Segment Tree

[![Build Status](https://travis-ci.org/fyibmsd/segment.svg?branch=master)](https://travis-ci.org/fyibmsd/segment)

### Intro
Segment tree is a tree data structure used for storing information about intervals, or segments. It allows querying which of the stored segments contain a given point.

### Benchmark

```sh
goos: darwin
goarch: amd64
pkg: github.com/fyibmsd/segment
BenchmarkQuerySum/query_sum_by_segment_tree-12         	20000000	       107 ns/op
BenchmarkQuerySum/query_sum_by_range-12                	10000000	       172 ns/op
```


[Wiki](https://en.wikipedia.org/wiki/Segment_tree)

