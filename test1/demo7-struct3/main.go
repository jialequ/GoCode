package main

import "fmt"

type A struct {
	Num int
}

// 方法，与类型绑定，属于结构体A，相当于类成员函数
func (a A) getNum() int {
	return a.Num
}

func main() {
	aa := A{}
	b := aa.getNum()
	fmt.Println(b)
}