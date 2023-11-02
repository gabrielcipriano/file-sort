package interfaces

type Writer interface {
	Write(record interface{}) error
}
