package core

import (
	"errors"
	"io"

	"github.com/gabrielcipriano/file-sort/interfaces"
)

type Joiner struct {
	left     FileStream
	right    FileStream
	writer   interfaces.Writer
	lessThan func(left interface{}, right interface{}) bool
}

func NewJoiner(
	left, right FileStream,
	writer interfaces.Writer,
	lessThan func(left interface{}, right interface{}) bool,
) *Joiner {
	return &Joiner{
		left, right, writer, lessThan,
	}
}

func (j *Joiner) Join() error {
	lVal, lErr := j.left.Get()
	if lErr != nil {
		return lErr
	}

	rVal, rErr := j.right.Get()
	if rErr != nil {
		return rErr
	}

	//while both streams are not empty
	for true {
		if j.lessThan(lVal, rVal) {
			lErr = j.writer.Write(lVal)
			if lErr != nil {
				return lErr
			}

			lVal, lErr = j.left.Get()
			if lErr != nil {
				if errors.Is(lErr, io.EOF) {
					break
				}
				return lErr
			}
		} else {
			rErr = j.writer.Write(rVal)
			if rErr != nil {
				return rErr
			}

			rVal, rErr = j.right.Get()
			if rErr != nil {
				if errors.Is(rErr, io.EOF) {
					break
				}
				return rErr
			}
		}
	}

	//while left stream is not empty
	for true {
		lErr = j.writer.Write(lVal)
		if lErr != nil {
			return lErr
		}

		lVal, lErr = j.left.Get()
		if lErr != nil {
			if errors.Is(lErr, io.EOF) {
				break
			}
			return lErr
		}
	}

	//while right stream is not empty
	for true {
		rErr = j.writer.Write(rVal)
		if rErr != nil {
			return rErr
		}

		rVal, rErr = j.right.Get()
		if rErr != nil {
			if errors.Is(rErr, io.EOF) {
				break
			}
			return rErr
		}
	}

	lErr = j.left.Finish()
	rErr = j.right.Finish()

	return nil
}
