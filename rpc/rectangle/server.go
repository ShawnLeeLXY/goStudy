package rectangle

import (
	"log"
	"net/http"
	"net/rpc"
)

/*
Go实现RPC程序，求矩形面积和周长
*/

type Params struct {
	Length, Width int
}

type Rect struct{}

// Area RPC服务端方法，求矩形面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Length * p.Width
	return nil
}

// Perimeter 周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Length + p.Width) * 2
	return nil
}

// server 服务端主函数
func server() {
	// 1. 注册一个Rect的服务
	rpc.Register(new(Rect))
	// 2. 服务处理绑定到http协议上
	rpc.HandleHTTP()
	// 3. 监听服务
	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		log.Panicln(err)
	}

	wg.Done()
}
