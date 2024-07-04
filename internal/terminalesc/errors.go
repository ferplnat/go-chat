package terminalesc

// FormatError is an error type that represents an invalid format specifier.
type FormatError struct {
	Err error
}

func (f *FormatError) Error() string {
	return "invalid format specifier"
}
