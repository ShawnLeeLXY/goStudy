package concurrent

import "fmt"

func chanDemo() {
	ch := make(chan int, 1) // 创建一个容量为1的有缓冲区通道
	ch <- 10                // main不会阻塞
	fmt.Println("发送成功")
}
