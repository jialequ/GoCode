package main

import "fmt"

type People interface {
    Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
    if think == "sb" {
        talk = "你是个大帅比"
    } else {
        talk = "您好"
    }
    return
}

func main() {
    var peo People = &Student{}
    var peo1 People = Student{}	// 报错, 非指针不能给指针
    think := "bitch"
    fmt.Println(peo.Speak(think))
}