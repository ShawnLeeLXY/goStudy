package netpack

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// tcp server端
func tcpServer() {
	ln, err := net.Listen("tcp", "127.0.0.1:8101")
	if err != nil {
		log.Fatal("listen failed, err:", err)
	}
	fmt.Printf("server正在监听...")
	for {
		conn, err := ln.Accept() // 建立连接
		if err != nil {
			log.Fatal("accept failed, err:", err)
		}
		go process(conn) // 启动一个goroutine处理连接
	}
}

func process(conn net.Conn) {
	defer conn.Close() // 关闭连接
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) // 读取数据
		if err != nil {
			log.Fatal("read from client failed, err:", err)
		}
		recv := string(buf[:n])
		fmt.Println("server收到了client发来的数据：", recv)
		conn.Write([]byte(recv)) // 发送数据
	}
}
