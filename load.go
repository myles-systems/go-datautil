// Copyright 2015 myles.systems. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package datautil

import (
	"errors"
	"io/ioutil"
	"net/http"
	"reflect"
)

// Load data from server or read a local file from disk.
// Multiple sources can be set and the Load func will return the data
// at the first success load.
func Load(src ...interface{}) ([]byte, error) {
	totalSources := len(src)
	if totalSources == 0 {
		return []byte{}, errors.New("Missing src parameter(s)")
	}

	var err error
	var data []byte
	for _, v := range src {
		switch reflect.TypeOf(v).Name() {
		case "Default":
			data = reflect.ValueOf(v).FieldByName("Data").Bytes()
			err = nil
			break
		case "string":
			data, err = loadSource(v.(string))
		}
		if err != nil {
			continue
		} else {
			break
		}
	}
	return data, err
}

func loadSource(src string) ([]byte, error) {
	// variable we return later
	var err error
	var data []byte
	// load or request the file/data
	switch GetSourceType(src) {
	case TypeFile:
		data, err = readFile(src)
		break
	case TypeURL:
		data, err = requestFile(src)
		break
	}
	return data, err
}

func readFile(src string) ([]byte, error) {
	data, err := ioutil.ReadFile(src)
	return data, err
}

func requestFile(src string) ([]byte, error) {
	res, err := http.Get(src)
	if err != nil {
		return []byte{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
