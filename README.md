# go-tprof

## Overview [![GoDoc](https://godoc.org/github.com/gokultp/go-tprof?status.svg)](https://godoc.org/github.com/gokultp/go-tprof) [![Build Status](https://travis-ci.org/gokultp/go-tprof.svg?branch=master)](https://travis-ci.org/gokultp/go-tprof) [![Code Climate](https://codeclimate.com/github/gokultp/go-tprof/badges/gpa.svg)](https://codeclimate.com/github/gokultp/go-tprof) [![Go Report Card](https://goreportcard.com/badge/github.com/gokultp/go-tprof)](https://goreportcard.com/report/github.com/gokultp/go-tprof)


[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/wRR9q7bJja0/0.jpg)](https://www.youtube.com/watch?v=wRR9q7bJja0)

## Prerequisites
    1. Go version >= 1.12 
    2. Node.js version >= 8.0.0 (for building the UI)
    3. Yarn
    4. GOPATH and local bin setup (`export PATH=$PATH:$GOPATH/bin`)
## Install


```
git clone https://github.com/gokultp/go-tprof.git
cd go-tprof
make config
make
```

## Usage Instructions

`tprof` is not a replacement for `go test`, it just parses the results and generates reports out of it.

To execute

```
    go test <pkg name> -v | tprof
```

Eg:
```
    go test encoding/... -v | tprof
```


## Features

- [x] Parser and reports logic.
- [x] HTML Page

## License

MIT.
