// package main

// import "fmt"

// // max
// // 形式参数
// // 实际参数
// func max(a, b int) int {
// 	var result int
// 	if a > b {
// 		result = a
// 	} else {
// 		result = b
// 	}
// 	return result
// }

// func getSum(nums ...int) int {
// 	sum := 0
// 	for _, num := range nums {
// 		sum += num
// 	}

// 	return sum
// }

// // 递归
// func getSumer(n int) int {
// 	if n == 1 {
// 		return 1
// 	}
// 	return getSumer(n-1) + n
// }

// // 切片，引用传递，地址一样
// func changeNum(slice []int) {
// 	fmt.Printf("%p\n", slice)
// }

// func main() {

// 	arr := [5]int{3, 2, 4, 1, 5}
// 	sli := arr[:4]
// 	fmt.Println(sli)
// }
