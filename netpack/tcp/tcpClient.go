package tcp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func TcpClient() {
	conn, err := net.Dial("tcp", "175.178.19.110:8101")
	if err != nil {
		log.Fatal("err:", err)
	}
	defer conn.Close() // 注册defer关闭连接
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("client已启动，请输入：")
	for {
		input, _ := reader.ReadString('\n') // 读取用户输入
		if input == "\n" || input == "\r" || input == "\r\n" {
			return
		}
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "QUIT" { // 如果输入quit就退出
			return
		}
		_, err = conn.Write([]byte(inputInfo)) // 发送数据
		if err != nil {
			log.Fatal("err:", err)
		}
		buf := [256]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Fatal("Recv failed, err:", err)
		}
		fmt.Println("client收到了server发来的数据：", string(buf[:n]))
	}
}
