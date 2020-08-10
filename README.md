# replacer 
![CI](https://github.com/guerinoni/replacer/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/guerinoni/replacer/branch/master/graph/badge.svg)](https://codecov.io/gh/guerinoni/replacer)
[![Go Report Card](https://goreportcard.com/badge/github.com/guerinoni/replacer)](https://goreportcard.com/report/github.com/guerinoni/replacer)

Command-line tool to rename a lot of files with some rules :)

# Feature
* Replace extension of files
* Change string contains in filenames

# Usage
```
replacer -h (help)
replacer -d . -ext txt c (change all file with extension "txt" to "c" in current dir )
replacer -d . -contains as 2 sa (change all file that contain "as" to "sa" in current dir)
```
