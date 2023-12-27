package main

import "fmt"

func main() {
	var str string = "hello"
	hmap := map[byte]byte{
		'h': 'h',
	}

	fmt.Println(hmap[str[0]])

	if hmap[str[0]] > 0 {
		fmt.Println("ok")
	} else {
		fmt.Println("no")
	}
}
