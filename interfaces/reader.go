package interfaces

type Reader interface {
	Read() (record interface{}, err error)
}
