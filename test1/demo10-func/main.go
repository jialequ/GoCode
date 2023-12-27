package main

import "fmt"

// 一个数从0开始，每次加上自己的值和当前循环次数, 然后*2, 这样迭代10次
// 普通实现
// func adder(x int) int {
//     return x * 2
// }

// func main() {
//     var a int
//     for i := 0; i < 10; i ++ {
// 		fmt.Println("a + i:", a + i)
//         a = adder(a + i)
//         fmt.Println("a", a)
//     }
// }

// 闭包实现
// 1. 外层函数返回内层函数
// 2. 内层函数使用了外层函数的变量
func adder() func(int) int {
    res := 0
    return func(x int) int {
        res = (res + x) * 2
        return res
    }
}

func main() {
    a := adder()
    for i := 0; i < 10; i++ {
        fmt.Println(a(i))
    }
}