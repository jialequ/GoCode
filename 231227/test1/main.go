package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := []int{8, 4, 6, 3, 5, 7, 9, 1}
	sort.Slice(s1, func(i, j int) bool { return s1[i] > s1[j] })
	fmt.Println(s1)
}
