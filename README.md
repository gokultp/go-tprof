# go-tprof

## Overview [![GoDoc](https://godoc.org/github.com/gokultp/go-tprof?status.svg)](https://godoc.org/github.com/gokultp/go-tprof) [![Build Status](https://travis-ci.org/gokultp/go-tprof.svg?branch=master)](https://travis-ci.org/gokultp/go-tprof) [![Code Climate](https://codeclimate.com/github/gokultp/go-tprof/badges/gpa.svg)](https://codeclimate.com/github/gokultp/go-tprof) [![Go Report Card](https://goreportcard.com/badge/github.com/gokultp/go-tprof)](https://goreportcard.com/report/github.com/gokultp/go-tprof)


## Install

1. Prerequisites

```
# should have go version >1.11
# shoud have a bin folder in your home directory (~/bin) and that should be part of PATH
export PATH=$PATH:~/bin/
```

```
git clone https://github.com/gokultp/go-tprof.git
cd go-tprof
go get -d
make 
```

## Usage Instructions

`tprof` is not a replacement for `go test`, it just parses the results and generates reports out of it.

To execute

```
    go test <pkg name> -v | tprof
```

## Features

- [x] Parser and reports logic.
- [] HTML Page

## License

MIT.