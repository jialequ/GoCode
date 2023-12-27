package main

import "fmt"

type Sayer interface {
	say()
}

type Dog struct{}
type Cat struct{}

func (d Dog) say() {
	fmt.Println("汪汪汪")
}

func (c Cat) say() {
	fmt.Println("喵喵喵")
}

func main() {
	d := Dog{}
	d.say()

	c := Cat{}
	c.say()

	var x Sayer
	x = Dog{}
	x.say()

	x = Cat{}
	x.say()
}