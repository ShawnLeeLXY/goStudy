package concurrent

import "fmt"

func receiver(ch chan int) {
	temp := <-ch
	fmt.Println("接收成功：", temp)
}

func syncDemo() {
	ch := make(chan int)
	go receiver(ch) // 启用goroutine从通道接收值
	num := 10
	ch <- num
	fmt.Println("发送成功：", num)
}
