package main

import "fmt"

// 类型断言 x.(T)
// x：表示类型为interface{}的变量
// T：表示断言x可能是的类型

func main() {
	var x interface{} = "pprof.cn"
	v, ok := x.(string)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("类型断言失败")
	}
}
