# go-datautil [![Build Status](https://travis-ci.org/myles-systems/go-datautil.svg?branch=master)](https://travis-ci.org/myles-systems/go-datautil) [![Coverage Status](https://coveralls.io/repos/myles-systems/go-datautil/badge.svg?branch=master&service=github)](https://coveralls.io/github/myles-systems/go-datautil?branch=master) [![GoDoc](https://godoc.org/github.com/myles-systems/go-datautil?status.svg)](https://godoc.org/github.com/myles-systems/go-datautil)

data utilities


## usage

```
import "github.com/myles-systems/go-datautil"

// load a local file
data, err := datautil.Load("path/to/file")

// load data from server (must be start with http)
data, err := datautil.Load("http://sample.tld/resource")

// load data with multiple paths (like fallbacks)
data, err := datautil.Load("http://sample.tld", "path/to/file", "other/fallback/file")

// load with defined default data
data, err := datautil.Load("path/to/file", udl.DefaultData("some data here"))
```
