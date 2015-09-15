package datautil

import (
	"net/url"
)

// Values used at the GetSourceType return integer.
const (
	TypeUnknown int = iota
	TypeFile
	TypeURL
)

// GetSourceType check the source type. is it an http(s)-url or a filepath?
func GetSourceType(src string) int {
	u, err := url.Parse(src)
	if err != nil {
		return TypeUnknown
	}

	tmpType := TypeUnknown
	switch u.Scheme {
	case "http":
		tmpType = TypeURL
		break
	case "https":
		tmpType = TypeURL
		break
	case "":
		if src != "" {
			tmpType = TypeFile
		}
		break
	}

	return tmpType
}
