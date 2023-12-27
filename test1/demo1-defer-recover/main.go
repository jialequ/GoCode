package main

import (
	"fmt"
)

func main() {
	test()
}

// 提高程序健壮性，defer+recover，捕获异常
func test() {
	// 利用defer+recover来捕获错误, defer后面加上匿名函数调用
	defer func() {
		// 调用recover捕获错误
		err := recover()
		// 如果没有错误, 返回值为nil
		if err != nil {
			// 错误已经捕获
			fmt.Println(err)
		}
	}()

	num1 := 10
	num2 := 0
	result := num1 / num2
	fmt.Println(result)
}