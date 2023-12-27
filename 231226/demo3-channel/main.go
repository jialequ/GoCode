package main

import (
	"fmt"
	"time"
)

func recv(c chan int) {
	ret := <-c
	fmt.Println(ret)
}
func main() {
	ch := make(chan int, 1)
	go recv(ch)
	ch <- 1
	// close(ch)
	time.Sleep(time.Second)
	fmt.Println("发送成功")
}
