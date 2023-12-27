package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	s1 := "100"
	i1, err := strconv.Atoi(s1)
	if err != nil {
		fmt.Println("conv err", err)
		return
	} else {
		fmt.Printf("type:%T value:%#v\n", i1, i1)
	}

	i2 := 200
	s2 := strconv.Itoa(i2)
	fmt.Println(reflect.TypeOf(s2))
	fmt.Println(s2)
}
