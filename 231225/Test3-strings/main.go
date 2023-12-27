package main

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

func main() {
	a := "hello"
	b := "world"

	// 1.普通操作
	// 字符串比较如果相等为0，不相等-1
	fmt.Println(strings.Compare(a, b))
	// 是否包含某字符或字符串
	fmt.Println(strings.Contains(a, "h"))
	// 任意一个字符存在都返回true 判断多个字符用&相连
	fmt.Println(strings.ContainsAny(a, "h&l"))
	// 字符串出现次数,如果后面的字符为空时返回字符串长度 +1
	fmt.Println(strings.Count(a, "o"))

	// 2.字符串分割
	// 切片可以使用fmt.Printf("%q",切片)打印
	s := "this is my world"
	// 用空格分割字符串, 返回[]string
	tmp1 := strings.Fields(s)
	fmt.Println(reflect.TypeOf(tmp1))
	// 指定字符分割
	strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace)
	s1 := "l,o,v,e"
	// 如果是空将分割每个字符
	tmp2 := strings.Split(s1, ",")
	fmt.Println(reflect.TypeOf(tmp2))
	// 分割后保留分割时的条件字符串
	fmt.Printf("%q\n", strings.SplitAfter(s1, ","))
	// n表示返回的切片的大小
	fmt.Printf("%q\n", strings.SplitN(s1, ",", 2))
	// n表示返回的切片的大小
	fmt.Printf("%q\n", strings.SplitAfterN(s1, ",", 2))

	// 3.字符出现的位置
	// 字符在字符串的位置 如果是中文字符可能是占两个字符
	var i int = strings.Index(s, "r")
	fmt.Println(i)

	// 4.反射包reflect.DeepEqual(want, got)可以比较切片是否相等
	got := strings.Split("a:b:c", ":") // 程序输出的结果
	want := []string{"a", "b", "c"}    // 期望的结果
	if reflect.DeepEqual(want, got) {
		fmt.Println("ok")
	}
}
