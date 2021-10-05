# replacer 
![CI](https://github.com/guerinoni/replacer/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/guerinoni/replacer/branch/master/graph/badge.svg)](https://codecov.io/gh/guerinoni/replacer)
[![lint](https://github.com/guerinoni/replacer/actions/workflows/lint.yml/badge.svg)](https://github.com/guerinoni/replacer/actions/workflows/lint.yml)
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
BenchmarkExecCamelCaseOneFile-8         	   14736	     71699 ns/op	    2280 B/op	      28 allocs/op
BenchmarkExecCamelCaseDir-8             	    1078	   1268006 ns/op	   22503 B/op	     422 allocs/op
BenchmarkExecCamelCaseLotDir-8          	      55	  20024731 ns/op	  611971 B/op	    4122 allocs/op
BenchmarkExecCamelCaseManyDir-8         	      15	  74275259 ns/op	 6993296 B/op	   10864 allocs/op
BenchmarkExecChangeContainsOneFile-8    	   14798	     82463 ns/op	    2264 B/op	      27 allocs/op
BenchmarkExecChangeContainsDir-8        	     956	   1163330 ns/op	   21177 B/op	     412 allocs/op
BenchmarkExecChangeContainsLotDir-8     	      68	  18553358 ns/op	  568569 B/op	    4024 allocs/op
BenchmarkExecChangeContainsManyDir-8    	      13	  81944776 ns/op	 6180158 B/op	   10684 allocs/op
BenchmarkExecChangeExtensionOneFile-8   	   16408	     69426 ns/op	    2264 B/op	      27 allocs/op
BenchmarkExecChangeExtensionDir-8       	    1114	   1763960 ns/op	   19175 B/op	     412 allocs/op
BenchmarkExecChangeExtensionLotDir-8    	      74	  13857763 ns/op	  389218 B/op	    4021 allocs/op
BenchmarkExecChangeExtensionManyDir-8   	       4	 341479902 ns/op	33262900 B/op	   42103 allocs/op
BenchmarkExecSnakeCaseOneFile-8         	   15810	     71889 ns/op	    2280 B/op	      28 allocs/op
BenchmarkExecSnakeCaseDir-8             	   88321	     14756 ns/op	    1931 B/op	      27 allocs/op
BenchmarkExecSnakeCaseLotDir-8          	     109	  10770991 ns/op	  611515 B/op	    4120 allocs/op
BenchmarkExecSnakeCaseManyDir-8         	      14	  83155920 ns/op	 5913929 B/op	   10600 allocs/op
PASS
ok  	github.com/guerinoni/replacer	24.268s
```
