package main

import (
	"fmt"
)

// func isRight(i *int) (a int, ok bool) {
// 	*i = 20
// 	return *i, true
// }

// func main() {
// 	i := 10
// 	i, ok := isRight(&i)
// 	if ok {
// 		fmt.Println("ok", i)
// 	} else {
// 		fmt.Println("no ok", i)
// 	}
// }

func main() {
	test()
}

func test() {
	var i *int

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err is:", err.(string))
		}
	}()

	fmt.Println(*i)
	fmt.Println("run success")

	a := 10
	b := 1
	fmt.Println(a + b)
}
