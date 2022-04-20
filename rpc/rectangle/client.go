package rectangle

import (
	"fmt"
	"log"
	"net/rpc"
)

// Client 客户端主函数
func client() {
	// 1. 连接远程rpc服务
	conn, err := rpc.DialHTTP("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal(err)
	}
	// 2. 调用方法
	for {
		fmt.Println("请输入矩形的长和宽：")
		var length, width int
		fmt.Scan(&length, &width)
		params := Params{length, width}
		ret := 0
		err2 := conn.Call("Rect.Area", params, &ret)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Println("面积=", ret)
		// 周长
		err3 := conn.Call("Rect.Perimeter", params, &ret)
		if err3 != nil {
			log.Fatal(err3)
		}
		fmt.Println("周长=", ret)
	}

	wg.Done()
}
