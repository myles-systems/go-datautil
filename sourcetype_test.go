package datautil

import (
	"testing"
)

func Test_GetSourceType_Unknown(t *testing.T) {
	sourceType := GetSourceType("")
	if sourceType != TypeUnknown {
		t.Error("Type was not checked as TypeUnknown")
	}
}

func Test_GetSourceType_File(t *testing.T) {
	sourceType := GetSourceType("path/to/file")
	if sourceType != TypeFile {
		t.Error("Type was not checked as TypeFile")
	}
}

func Test_GetSourceType_URL(t *testing.T) {
	sourceType := GetSourceType("http://localhost:8000")
	if sourceType != TypeURL {
		t.Error("Type was not checked as TypeURL (http)")
	}
	sourceType = GetSourceType("https://localhost:8000")
	if sourceType != TypeURL {
		t.Error("Type was not checked as TypeURL (https)")
	}
}
