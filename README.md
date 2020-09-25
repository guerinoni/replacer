# replacer 
![CI](https://github.com/guerinoni/replacer/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/guerinoni/replacer/branch/master/graph/badge.svg)](https://codecov.io/gh/guerinoni/replacer)
[![Go Report Card](https://goreportcard.com/badge/github.com/guerinoni/replacer)](https://goreportcard.com/report/github.com/guerinoni/replacer)

Command-line tool to rename a lot of files with some rules :)

# Feature

- [x] help  `replacer -h`
- [x] change extension `replacer -d . -ext txt c` (-d = directory)
- [x] change contains `replacer -d . -contains as ss`
- [x] convert to snake_case `replacer -snake <camelCaseFile>` or `replacer -snake <folder>`
- [ ] convert to camelCase
- [ ] remane list of files with incremental index