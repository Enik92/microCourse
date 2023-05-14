package src

import (
	"io"
	"unicode"
)

type Reader interface {
	Read(p []byte) (int, error)
	ReadAll(bufSize int) (string, error)
	BytesRead() int64
}

type CountingToLowerReaderImpl struct {
	Reader         io.Reader
	TotalBytesRead int64
}

func (cr *CountingToLowerReaderImpl) Read(p []byte) (n int, err error) {
	n, err = cr.Reader.Read(p)
	if err != nil {
		return 0, err
	}

	for i, v := range p {
		v = byte(unicode.ToLower(rune(p[i])))
		p[i] = v
	}

	cr.TotalBytesRead += int64(n)

	return n, err
}

func (cr *CountingToLowerReaderImpl) ReadAll(bufSize int) (str string, err error) {
	b := make([]byte, bufSize)

	n, err := cr.Read(b)
	for ; err == nil; n, err = cr.Read(b) {
		str += string(b[:n])
	}

	if err == io.EOF {
		return str, nil
	}

	return
}

func (cr *CountingToLowerReaderImpl) BytesRead() int64 {
	b := cr.TotalBytesRead

	return b
}

func NewCountingReader(r io.Reader) *CountingToLowerReaderImpl {
	return &CountingToLowerReaderImpl{
		Reader: r,
	}
}
