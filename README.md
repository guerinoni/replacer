# replacer 
![CI](https://github.com/guerinoni/replacer/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/guerinoni/replacer/branch/master/graph/badge.svg)](https://codecov.io/gh/guerinoni/replacer)
[![Go Report Card](https://goreportcard.com/badge/github.com/guerinoni/replacer)](https://goreportcard.com/report/github.com/guerinoni/replacer)

Command-line tool to rename a lot of files with some rules :)

# Feature

- [x] help  `replacer -h`
- [x] change extension `replacer -d . -ext txt c` (-d = directory)
- [x] change contains `replacer -d . -contains as ss`
- [x] convert to snake_case `replacer -d . -snake`
- [x] convert to camelCase `replacer -d . -camel`
- [ ] rename list of files with incremental index


## Benchmark

```
go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/guerinoni/replacer
cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
BenchmarkExecCamelCaseOneFile-8         	   43515	     27391 ns/op	    2392 B/op	      52 allocs/op
BenchmarkExecCamelCaseDir-8             	    2485	    480512 ns/op	  191411 B/op	     692 allocs/op
BenchmarkExecCamelCaseLotDir-8          	     186	   5843872 ns/op	 2327337 B/op	    6812 allocs/op
BenchmarkExecCamelCaseManyDir-8         	      30	  35512414 ns/op	10069383 B/op	   11014 allocs/op
BenchmarkExecChangeContainsOneFile-8    	   45991	     26119 ns/op	    2168 B/op	      25 allocs/op
BenchmarkExecChangeContainsDir-8        	    2462	    543166 ns/op	  188296 B/op	     422 allocs/op
BenchmarkExecChangeContainsLotDir-8     	     195	   5716483 ns/op	 2297134 B/op	    4112 allocs/op
BenchmarkExecChangeContainsManyDir-8    	      32	  36024406 ns/op	10076980 B/op	   11033 allocs/op
BenchmarkExecChangeExtensionOneFile-8   	   47062	     26028 ns/op	    2168 B/op	      25 allocs/op
BenchmarkExecChangeExtensionDir-8       	    2478	    497402 ns/op	  185996 B/op	     422 allocs/op
BenchmarkExecChangeExtensionLotDir-8    	     207	   5899122 ns/op	 2088637 B/op	    4112 allocs/op
BenchmarkExecChangeExtensionManyDir-8   	       6	 177506436 ns/op	45918349 B/op	   41012 allocs/op
BenchmarkExecSnakeCaseOneFile-8         	   40345	     27524 ns/op	    2408 B/op	      53 allocs/op
BenchmarkExecSnakeCaseDir-8             	  230889	      6525 ns/op	    1602 B/op	      20 allocs/op
BenchmarkExecSnakeCaseLotDir-8          	     216	   5691752 ns/op	 2327114 B/op	    6912 allocs/op
BenchmarkExecSnakeCaseManyDir-8         	      31	  38836712 ns/op	10068742 B/op	   11014 allocs/op
PASS
ok  	github.com/guerinoni/replacer	24.268s
```
