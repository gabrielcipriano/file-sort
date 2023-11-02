package core

import (
	"errors"
	"io"

	"github.com/gabrielcipriano/file-sort/interfaces"
)

var (
	ErrInvalidReader = errors.New("Invalid Reader")
	ErrInvalidSize   = errors.New("Size not valid")
)

type FileStream struct {
	chunck   *CircularChunck
	r        interfaces.Reader
	onFinish func() error
}

func NewFileStream(
	reader interfaces.Reader, size int,
	fnFinish func() error,
) (*FileStream, error) {
	if reader == nil {
		return nil, ErrInvalidReader
	}
	if size <= 0 {
		return nil, ErrInvalidSize
	}
	return &FileStream{
		chunck:   NewCircularChunck(size),
		r:        reader,
		onFinish: fnFinish,
	}, nil
}

func (fs FileStream) Get() (interface{}, error) {
	for !fs.chunck.Full() {
		elem, err := fs.r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
		ok := fs.chunck.Push(elem)
		if !ok {
			return nil, errors.New("Can not push a new element!")
		}
	}
	elem, ok := fs.chunck.Pop()
	if !ok {
		return nil, io.EOF
	}
	return elem, nil
}

func (fs FileStream) Finish() error {
	return fs.onFinish()
}
