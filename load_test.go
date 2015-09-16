// Copyright 2015 myles.systems. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package datautil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Load_File(t *testing.T) {
	data, err := Load("fixture/test.txt")
	if err != nil {
		t.Error(err)
	}
	if string(data) != "hello world\n" {
		t.Error("File data not equal")
	}
}

func Test_Load_URL(t *testing.T) {
	testsrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello world")
	}))
	defer testsrv.Close()

	data, err := Load(testsrv.URL)
	if err != nil {
		t.Error(err)
	}
	if string(data) != "hello world\n" {
		t.Error("File data not equal")
	}
}

func Test_Load_MultipleSources(t *testing.T) {
	data, err := Load("fixture/foo.txt", "fixture/test.txt", "fixture/bar.txt", DefaultString("def"))
	if err != nil {
		t.Error(err)
	}
	if string(data) != "hello world\n" {
		t.Error("File data not equal")
	}
}

func Test_Load_MultipleSourcesAndDefault(t *testing.T) {
	data, err := Load("fixture/foo.txt", "fixture/bar.txt", DefaultString("def"))
	if err != nil {
		t.Error(err)
	}
	if string(data) != "def" {
		t.Error("File data not equal")
	}
}
