package core

type CircularChunck struct {
	data     []interface{}
	start    int
	size     int
	capacity int
	less     func(left interface{}, right interface{}) bool
}

func NewCircularChunck(size int) *CircularChunck {
	return &CircularChunck{
		data:     make([]interface{}, size),
		start:    0,
		size:     0,
		capacity: size,
	}
}

func (cc *CircularChunck) Push(elem interface{}) bool {
	if cc.size == cc.capacity {
		return false
	}
	cc.data[(cc.start+cc.size)%cc.capacity] = elem
	cc.size++
	return true
}

func (cc *CircularChunck) Pop() (interface{}, bool) {
	top, ok := cc.top()
	if ok {
		cc.start = (cc.start + 1) % cc.capacity
		cc.size--
	}
	return top, ok
}

func (cc *CircularChunck) top() (interface{}, bool) {
	if cc.size == 0 {
		return nil, false
	}
	return cc.data[cc.start], true
}

func (cc *CircularChunck) Clear() {
	cc.size = 0
	cc.start = 0
}

func (cc *CircularChunck) Full() bool {
	return cc.size == cc.capacity
}

func (cc *CircularChunck) Empty() bool {
	return cc.size == 0
}

func (cc *CircularChunck) Len() int { return cc.size }
func (cc *CircularChunck) Swap(i, j int) {
	left := (i + cc.start) % cc.capacity
	right := (j + cc.start) % cc.capacity
	cc.data[left], cc.data[right] = cc.data[right], cc.data[left]
}
func (cc *CircularChunck) Less(i, j int) bool {
	left := (i + cc.start) % cc.capacity
	right := (j + cc.start) % cc.capacity
	return cc.less(cc.data[left], cc.data[right])
}
