package main

import "fmt"

// 定义结构体，放入人的属性
type Person struct {
	// 变量名字大写，外界才可以访问
	Name string
	Age int
	Company string
}

// func main() {
// 	// 创建结构体实例,直接创建
// 	var mine Person
// 	fmt.Println(mine)
// 	mine.Name = "屈佳乐"
// 	mine.Age = 23
// 	mine.Company = "烽火星空"
// 	fmt.Println(mine)
// }

// func main() {
// 	// 创建时初始化
// 	var m Person = Person{"屈佳乐", 23, "烽火星空"}
// 	fmt.Println(m)
// }

// func main() {
// 	var m *Person = new(Person)
// 	(*m).Age = 23
// 	// 底层优化，可以直接用m操作
// 	m.Company = "烽火"
// 	fmt.Println(*m)
// }

func main() {
	var m *Person = &Person{}
	m.Age = 23
	fmt.Println(*m)
}