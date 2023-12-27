package main

import "fmt"

type Stu struct {
	Name string
	Age int8
}

func newStudent(name string, age int8) *Stu {
	ts := &Stu{name, age}
	return ts
}

func (s Stu) printInfo(ret int8) {
	fmt.Println(s.Name, "现在的年龄是:", s.Age)
	age := s.Age + ret
	fmt.Println(s.Name, ret, "年后的年龄是:", age)
}

func (s *Stu) setAge(ret int8) {
	s.Age = ret
}

func main() {
	s1 := newStudent("屈佳乐", 23)
	s1.printInfo(10)
	s1.setAge(40)
	fmt.Println(*s1)
}