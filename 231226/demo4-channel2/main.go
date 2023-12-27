package main

import "fmt"

// 匿名函数管道写，main管道读
func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	for i := 0; i < 5; i++ {
		if v, ok := <-ch; ok {
			fmt.Println(v)
		}
	}

	fmt.Println("over")
}
