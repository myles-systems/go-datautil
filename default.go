package datautil

type Default struct {
	Data []byte
}

func DefaultString(d string) Default {
	tmp := []byte(d)
	return Default{tmp}
}

func DefaultBytes(d []byte) Default {
	return Default{d}
}
