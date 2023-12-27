package main

import "fmt"

// 单值接收
// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		ch <- "hello"
// 	}()

// 	str := <-ch

// 	fmt.Println(str)
// }

func main() {
    ch := make(chan int)

    go func() {
        for i := 0; i < 5; i++ {
			// 发送数据到通道
            ch <- i 
        }

		// 关闭通道，关闭后只能接收，不能发送
        close(ch) 
    }()

    for value := range ch { // 从通道循环接收数据
        fmt.Println(value)
    }
}