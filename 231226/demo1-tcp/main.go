package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 延迟关闭连接
	defer conn.Close()

}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1")
	if err != nil {
		fmt.Println("listen error, err:", err)
		return
	}

	for {
		// 获取连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error, err:", err)
			continue
		}

		go process(conn)
	}
}
