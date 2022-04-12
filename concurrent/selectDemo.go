package concurrent

import (
	"fmt"
	"math/rand"
	"time"
)

func selectDemo() {
	chSlice := make([]chan string, 5)
	for i := 0; i < 5; i++ {
		chSlice[i] = make(chan string)
		ind := i
		go func(ind int) {
			time.Sleep(time.Millisecond * (time.Duration)(rand.Int()%100))
			chSlice[ind] <- fmt.Sprintf("channel %d say: hello!\n", ind)
		}(ind)
	}
	// select同时监听多个chan 直到其中一个ready
	select {
	case s0 := <-chSlice[0]:
		fmt.Println(s0)
	case s1 := <-chSlice[1]:
		fmt.Println(s1)
	case s2 := <-chSlice[2]:
		fmt.Println(s2)
	case s3 := <-chSlice[3]:
		fmt.Println(s3)
	case s4 := <-chSlice[4]:
		fmt.Println(s4)
	}
}
