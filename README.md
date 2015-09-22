# go-datautil [![Build Status](https://travis-ci.org/myles-systems/go-datautil.svg?branch=master)](https://travis-ci.org/myles-systems/go-datautil) [![Coverage Status](https://coveralls.io/repos/myles-systems/go-datautil/badge.svg?branch=master&service=github)](https://coveralls.io/github/myles-systems/go-datautil?branch=master) [![GoDoc](https://godoc.org/github.com/myles-systems/go-datautil?status.svg)](https://godoc.org/github.com/myles-systems/go-datautil)

**go-datautil** makes it easy to load data from a local file or URL (with fallbacks). 


## usage

```
import "github.com/myles-systems/go-datautil"

// load data from local file
data, err := datautil.Load("path/to/file")

// load data from server
data, err := datautil.Load("http://sample.tld/resource")

// load data from server (with fallbacks)
data, err := datautil.Load("http://sample.tld", "path/to/file", "other/fallback/file")

// load default data if local file is inaccessible
data, err := datautil.Load("path/to/file", udl.DefaultData("some data here"))
```
