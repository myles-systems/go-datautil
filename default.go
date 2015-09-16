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
