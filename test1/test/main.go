package main

import "fmt"

// func main() {
//     s := []int{1, 2, 3, 4, 5}
// 	fmt.Printf("1:%p\n", &s)
// 	fmt.Printf("1:%p\n", &s[0])
// 	fmt.Printf("1:%p\n", &s[1])
// 	fmt.Printf("1:%p\n", &s[2])
// 	fmt.Printf("1:%p\n", &s[3])
// 	fmt.Printf("1:%p\n", &s[4])

// 	// 复制 struct slice { pointer, len, cap }
//     for i, v := range s {
//         if i == 0 {
// 			// 对 slice 的修改，不会影响 range
//             s = s[:3]
//         	fmt.Println(s, "len(s):", len(s))
//         	fmt.Printf("2:%p\n", &s)
// 			// 对底层数据的修改
//             s[2] = 100
// 			fmt.Printf("3:%p\n", &s[2])
//         }
// 		fmt.Printf("4:%p\n", &v)
//         fmt.Println(i, v)
//     }
// }

func calCulate(a, b int, fun func(int, int) int) int {
	return fun(a, b)
}

func main() {
	a := 20
	b := 10
	c := calCulate(a, b, func(x, y int) int {return x + y})
	fmt.Println(c)
	c = calCulate(a, b, func(x, y int) int {return x - y})
	fmt.Println(c)
	c = calCulate(a, b, func(x, y int) int {return x * y})
	fmt.Println(c)
	c = calCulate(a, b, func(x, y int) int {return x / y})
	fmt.Println(c)
}
