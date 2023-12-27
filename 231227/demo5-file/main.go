package main

import (
	"fmt"
	"os"
)

// 写文件
func main() {
	// 新建文件
	file, err := os.Create("./1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for i := 0; i < 5; i++ {
		file.WriteString("ab\n")
		file.Write([]byte("cd\n"))
		file.Write([]byte("1234\n"))
	}
}

// 读文件
// func main() {
// 	// 打开文件
// 	file, err := os.Open("./1.txt")
// 	if err != nil {
// 		fmt.Println("open file err :", err)
// 		return
// 	}

// 	defer file.Close()
// 	// 定义接收文件读取的字节数组
// 	var buf [128]byte
// 	var content []byte
// 	for {
// 		n, err := file.Read(buf[:])
// 		// io.EOF文件结尾
// 		if err == io.EOF {
// 			// 读取结束
// 			break
// 		}
// 		if err != nil {
// 			fmt.Println("read file err ", err)
// 			return
// 		}
// 		content = append(content, buf[:n]...)
// 	}
// 	fmt.Println(string(content))
// }
