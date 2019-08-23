# go-tprof

## Overview [![GoDoc](https://godoc.org/github.com/gokultp/go-tprof?status.svg)](https://godoc.org/github.com/gokultp/go-tprof) [![Build Status](https://travis-ci.org/gokultp/go-tprof.svg?branch=master)](https://travis-ci.org/gokultp/go-tprof) [![Code Climate](https://codeclimate.com/github/gokultp/go-tprof/badges/gpa.svg)](https://codeclimate.com/github/gokultp/go-tprof) [![Go Report Card](https://goreportcard.com/badge/github.com/gokultp/go-tprof)](https://goreportcard.com/report/github.com/gokultp/go-tprof)


## Install

```
go get github.com/gokultp/go-tprof/cmd/tprof
go install github.com/gokultp/go-tprof/cmd/tprof 
```

## Usage Instructions

`tprof` is not a replacement for `go test`, it just parses the results and generates reports out of it.

To execute

```
    go test <pkg name> -v | tprof
```

## Features

- [x] Parser and reports logic.
- [ ] HTML Page

## License

MIT.
