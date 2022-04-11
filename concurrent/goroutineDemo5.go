package concurrent

import (
	"fmt"
	"runtime"
	"time"
)

func printA() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func printB() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func goroutineDemo5() {
	// 设置执行核心数量
	runtime.GOMAXPROCS(2)
	go printA()
	go printB()
	time.Sleep(time.Second)
}
