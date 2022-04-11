package concurrent

import "fmt"

// 打印整数0~99的平方
func chanDemo3() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine 将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine 从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			// 若通道关闭后再取值 ok=false
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值并打印
	// 通道关闭后会退出for range循环
	for i := range ch2 {
		fmt.Println(i)
	}
}
