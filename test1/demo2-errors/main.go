package main

import (
	"fmt"
	"errors"
)

func main() {
	err := test()
	if err != nil {
		panic(err)
	}

	fmt.Println("除法执行成功~")
}

// 自定义错误，使用errors包下面的new函数
func test() (err error) {
	num1 := 10
	num2 := 0
	if num2 == 0 {
		// 抛出自定义错误
		return errors.New("分母不能为0~")
	} else {
		result := num1 / num2
		fmt.Println(result)
		
		//如果没有错误，返回nil
		return nil
	}
}