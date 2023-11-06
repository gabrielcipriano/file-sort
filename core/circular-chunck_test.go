package core_test

import (
	"errors"

	core "github.com/gabrielcipriano/file-sort/core"
)

var (
	errCantPush = errors.New("Can't Push")
	errCantPop  = errors.New("Can't Pop")
)

func addElements(cc *core.CircularChunck, n int) (int, error) {
	for i := 0; i < n; i++ {
		if !cc.Push(i) {
			return i, errCantPush
		}
	}
	return n, nil
}

func removeElements(cc *core.CircularChunck, n int) ([]int, error) {
	answer := make([]int, 0, n)
	for i := 0; i < n; i++ {
		elem, ok := cc.Pop()
		if !ok {
			return answer, errCantPop
		}
		answer = append(answer, elem.(int))
	}
	return answer, nil
}

// func main() {
// 	size := 10
// 	cc := core.NewCircularChunck(size)

// 	_, err := addElements(cc, size)
// 	if err != nil {
// 		panic(err)
// 	}

// 	_, err := removeElements(cc, size)
// 	if err != nil {
// 		panic(err)
// 	}

// }
