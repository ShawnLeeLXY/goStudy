package concurrent

import (
	"fmt"
	"sync"
)

func mutexDemo() {
	var num int
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go add(&num, &wg, &lock)
	}
	wg.Wait()
	fmt.Println(num) //3000
}

func add(num *int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*num += 1
		lock.Unlock()
	}
	wg.Done()
}
