package concurrent

import (
	"fmt"
	"sync"
)

// 一个简单的多线程同步示例
func goroutineDemo1() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}

var wg sync.WaitGroup
