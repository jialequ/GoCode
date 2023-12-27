package main

import (
	"fmt"
	"log"
	"os"
)

// 标准的logger通常会把写日志的配置写在init函数中
func init() {
	logFile, err := os.OpenFile("./demo3.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	// 设置输出文件，前提是要先打开文件
	log.SetOutput(logFile)
	// 设置输出选项
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func main() {
	// 设置输出选项
	// log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	// log.Println("这是一条普通的日志")

	// // 设置输出前缀
	// log.SetPrefix("[jlqu]")
	// name := "jialequ"
	// log.Printf("这是一条%s日志\n", name)

	// 设置logger的输出目的地，默认是标准错误输出
	// 标准的logger通常会把写日志的配置写在init函数中
	log.Println("这是一条很普通的日志。")
	// 设置输出前缀
	log.SetPrefix("[小王子]")
	log.Println("这是一条很普通的日志。")
}
