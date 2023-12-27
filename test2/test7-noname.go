package main

// 匿名函数

// func main() {

// 	func() {
// 		fmt.Println("我是f3函数")
// 	}()

// 	// 匿名函数自己调用自己，可以传参
// 	ret := func(a, b int) int {
// 		return a + b
// 	}(1, 2)

// 	fmt.Println(ret)
// }

// 闭包
func main() {
	f1 := increment()
	f1()
}

func increment() func() int {
	i := 0
	fun := func() int {
		// 定义了匿名函数但是没有调用
		i++
		return i
	}

	return fun
}
