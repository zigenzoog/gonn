package gonn

// ReadWriter
type ReadWriter interface {
	Reader
	Writer
}

// Reader
type Reader interface {
	Read(Reader) error
}

// Writer
type Writer interface {
	Write(...Writer) error
}

// Filer
type Filer interface {
	ReadWriter
	GetValue(key string) interface{}
	FileName() (string, error)
}