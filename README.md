# replacer 
![CI](https://github.com/guerinoni/replacer/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/guerinoni/replacer/branch/master/graph/badge.svg)](https://codecov.io/gh/guerinoni/replacer)
[![Go Report Card](https://goreportcard.com/badge/github.com/guerinoni/replacer)](https://goreportcard.com/report/github.com/guerinoni/replacer)

Command-line tool to rename a lot of files with some rules :)

# Feature

- [x] help  `replacer -h`
- [x] change extension `replacer -d . -ext txt c` (-d = directory)
- [x] change contains `replacer -d . -contains as ss`
- [x] convert to snake_case `replacer -d -snake`
- [ ] convert to camelCase
- [ ] rename list of files with incremental index


## Benchmark

```
BenchmarkExecChangeContainsOneFile    	    1000000000	         0.000039 ns/op
BenchmarkExecChangeContainsDir        	    1000000000	         0.000643 ns/op
BenchmarkExecChangeContainsLotDir     	    1000000000	         0.00601 ns/op
BenchmarkExecChangeContainsManyDir    	    1000000000	         0.0446 ns/op

BenchmarkExecChangeExtensionOneFile   	    1000000000	         0.000040 ns/op
BenchmarkExecChangeExtensionDir       	    1000000000	         0.000547 ns/op
BenchmarkExecChangeExtensionLotDir    	    1000000000	         0.00609 ns/op
BenchmarkExecChangeExtensionManyDir   	    1000000000	         0.215 ns/op

BenchmarkExecSnakeCaseOneFile         	    1000000000	         0.000055 ns/op
BenchmarkExecSnakeCaseDir             	    1000000000	         0.000628 ns/op
BenchmarkExecSnakeCaseLotDir          	    1000000000	         0.00645 ns/op
BenchmarkExecSnakeCaseManyDir         	    1000000000	         0.0408 ns/op
```
