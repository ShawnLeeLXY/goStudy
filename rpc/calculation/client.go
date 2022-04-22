package calculation

import (
	"fmt"
	"log"
	"net/rpc"
)

type ArithReq struct {
	A, B int
}

type ArithResp struct {
	Pro, Quo, Rem int
}

func client() {
	conn, err := rpc.DialHTTP("tcp", "175.178.19.110:8101")
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := ArithReq{}
		resp := ArithResp{}
		fmt.Println("请输入A和B的值：")
		fmt.Scan(&req.A, &req.B)
		err := conn.Call("Arith.Multiply", req, &resp)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d和%d的乘积是%d\n", req.A, req.B, resp.Pro)
		err2 := conn.Call("Arith.Divide", req, &resp)
		if err2 != nil {
			log.Fatal(err2)
		}
		fmt.Printf("%d和%d的商是%d，余数是%d\n", req.A, req.B, resp.Quo, resp.Rem)
	}
}
