package main

import "fmt"

func main() {
	// hmap := map[int]int{}
	hmap := make(map[int]int)
	// 增加元素
	hmap[10] = 5
	hmap[3] = 6
	hmap[4] = 7
	fmt.Println(hmap, len(hmap))

	// 查找元素
	_, ok := hmap[4]
	if ok {
		// 修改
		hmap[4] = 4
		fmt.Println(hmap, len(hmap))
	} else {
		fmt.Println("not find it")
	}

	// 清空hmap
	for k := range hmap {
		delete(hmap, k)
	}
	fmt.Println(hmap, len(hmap))
}