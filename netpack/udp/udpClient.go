package udp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func UdpClient() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8101,
	})
	if err != nil {
		fmt.Println("连接服务端失败，err:", err)
		return
	}
	defer socket.Close()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("client已启动，请输入：")
	for {
		input, _, _ := reader.ReadLine()
		_, err = socket.Write(input) // 发送数据
		if err != nil {
			log.Fatal("client发送数据失败，err:", err)
		}
		buf := make([]byte, 256)
		n, addr, err := socket.ReadFromUDP(buf) // 接收数据
		if err != nil {
			log.Fatal("client接收数据失败，err:", err)
		}
		fmt.Printf("client收到数据：buf:%v addr:%v count:%v\n", string(buf[:n]), addr, n)
	}
}
