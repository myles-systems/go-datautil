// Copyright 2015 myles.systems. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package datautil

// Default store the default data for a Load call.
type Default struct {
	Data []byte
}

// DefaultString set a string to a Default struct and return it.
func DefaultString(d string) Default {
	tmp := []byte(d)
	return Default{tmp}
}

// DefaultBytes set a byte array to a Default struct and return it.
func DefaultBytes(d []byte) Default {
	return Default{d}
}
