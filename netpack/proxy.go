package netpack

func Proxy() {
	go tcpServer()
	tcpClient()
}
