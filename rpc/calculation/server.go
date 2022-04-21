package calculation

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//type ArithReq struct {
//	A, B int
//}
//
//type ArithResp struct {
//	Pro, Quo, Rem int
//}

type Arith struct{}

func (a *Arith) Multiply(req ArithReq, resp *ArithResp) error {
	resp.Pro = req.A * req.B
	return nil
}

func (a *Arith) Divide(req ArithReq, resp *ArithResp) error {
	if req.B == 0 {
		return errors.New("Divided by zero!")
	}
	resp.Quo = req.A / req.B
	resp.Rem = req.A % req.B
	return nil
}

func server() {
	rpc.Register(new(Arith))
	rpc.HandleHTTP()
	ln, err := net.Listen("tcp", ":8989")
	if err != nil {
		log.Fatal("Arith error:", err)
	}
	go http.Serve(ln, nil)
}
