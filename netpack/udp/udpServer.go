package udp

import (
	"fmt"
	"log"
	"net"
)

func UdpServer() {
	ln, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8101,
	})
	if err != nil {
		log.Fatal("Listen failed, err:", err)
	}
	defer ln.Close()
	fmt.Println("server正在监听...")
	for {
		var buf [256]byte
		n, addr, err := ln.ReadFromUDP(buf[:]) // 接收数据
		if err != nil {
			log.Println("Read udp failed, err:", err)
		}
		fmt.Printf("server收到数据：buf:%v addr:%v count:%v\n", string(buf[:n]), addr, n)
		_, err = ln.WriteToUDP(buf[:n], addr) // 发送数据
		if err != nil {
			log.Println("Write to udp failed, err:", err)
		}
	}
}
