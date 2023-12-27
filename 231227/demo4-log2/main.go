package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// 日志文件
var logFile io.Writer

// 自定义的logger
var logger *log.Logger

func init() {
	logtmp, err := os.OpenFile("./demo4.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	logFile = logtmp
	if err != nil {
		fmt.Println("open log file error", err)
		return
	}

	logger = log.New(logFile, "[demo4]", log.Llongfile|log.Ldate|log.Ltime)
}

func main() {
	logger.Println("这是一条测试的log")
}
