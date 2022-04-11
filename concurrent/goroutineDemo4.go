package concurrent

import (
	"fmt"
	"runtime"
	"time"
)

func goroutineDemo4() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			// 后续代码不执行
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	// main()函数5s后退出
	time.Sleep(time.Second * 5)
}
