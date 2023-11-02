// func main() {
// 	size := 10
// 	cc := core.NewCircularChunck(size)

// 	for i := 0; i < size; i++ {
// 		err := cc.Push(i)
// 		if err != nil {
// 			panic(err)
// 		}
// 	}

// 	for i := 0; i < size; i++ {
// 		val, err := cc.Pop()
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(val)
// 	}

// 	val, err := cc.Pop()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(val)

// }