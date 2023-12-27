package main

import "fmt"

type Student struct{
	Age int
}

type Person struct {
	Age int
}

// 取别名的类型也需要强转
type stu Student

func main() {
	// 结构体转换必须拥有完全相同的字段
	p := Person{10}
	s := Student{10}
	s = Student(p)
	fmt.Println(p)
	fmt.Println(s)
}