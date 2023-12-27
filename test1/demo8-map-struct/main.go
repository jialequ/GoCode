package main

import "fmt"

type Stu struct {
	Name string
	Age int8
}

func main() {
	// map，key : 名字， value：属性结构体
	m := make(map[string]*Stu)
	arr := []Stu{
		{"屈佳乐", 23},
		{"赵佳林", 24},
	}

	// m赋值
	for _, v := range arr {
		m[v.Name] = &v
	}
	
	for k, v := range m {
		fmt.Println(k, *v)
	}
}