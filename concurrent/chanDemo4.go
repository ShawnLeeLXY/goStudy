package concurrent

import "fmt"

// outCh是一个单向通道 对于goroutine来说只能发送
func counter(outCh chan<- int) {
	for i := 0; i < 100; i++ {
		outCh <- i
	}
	close(outCh)
}

// inCh是一个单向通道 对于goroutine来说只能接收
func squarer(outCh chan<- int, inCh <-chan int) {
	for {
		// 若通道关闭后再取值 ok==false
		i, ok := <-inCh
		if !ok {
			break
		}
		outCh <- i * i
	}
	close(outCh)
}

func printer(inCh <-chan int) {
	// 通道关闭后会退出for range循环
	for i := range inCh {
		fmt.Println(i)
	}
}

// 打印整数0~99的平方
func chanDemo4() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
