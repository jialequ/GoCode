package main

import "fmt"

func main() {
	var scores [5]int

	// 输入成绩
	for i := 0; i < len(scores); i++ {
		fmt.Scanf("%d", &scores[i])
	}

	// 求总分和平均分
	sum := 0
	for _, score := range scores {
		sum += score
	}

	age := sum / len(scores)

	fmt.Println("sum:", sum, "age:", age)
}