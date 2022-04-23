package http

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HttpClient() {
	resp, _ := http.Get("http://127.0.0.1:8101/go")
	defer resp.Body.Close()
	// 200 OK
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)

	buf := make([]byte, 1024)
	for {
		// 接收服务端信息
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else {
			fmt.Println("client读取完毕")
			fmt.Println(string(buf[:n]))
			break
		}
	}
}
