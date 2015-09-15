package datautil

import (
	"testing"
)

func Test_DefaultString(t *testing.T) {
	tmp := DefaultString("hello string")
	if string(tmp.Data) != "hello string" {
		t.Error("DefaultString func failed")
	}
}

func Test_DefaultBytes(t *testing.T) {
	tmpBytes := []byte(`hello bytes`)
	tmp := DefaultBytes(tmpBytes)
	if string(tmp.Data) != "hello bytes" {
		t.Error("DefaultBytes func failed")
	}
}
