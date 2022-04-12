package concurrent

import (
	"fmt"
	"sync"
	"time"
)

func RWMutexDemo() {
	var wg sync.WaitGroup
	var rwLock sync.RWMutex
	start := time.Now()
	wg.Add(1012)
	go func() {
		for i := 0; i < 1000; i++ {
			go read(&wg, &rwLock)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10; i++ {
			go write(&wg, &rwLock)
		}
		wg.Done()
	}()
	wg.Wait()
	end := time.Now()
	fmt.Println("total time:", end.Sub(start))
}

func read(wg *sync.WaitGroup, rwLock *sync.RWMutex) {
	rwLock.RLock()
	time.Sleep(time.Millisecond) // 假设读操作用时1ms
	rwLock.RUnlock()
	wg.Done()
}

func write(wg *sync.WaitGroup, rwLock *sync.RWMutex) {
	rwLock.Lock()
	time.Sleep(time.Millisecond * 100) // 假设写操作用时100ms
	rwLock.Unlock()
	wg.Done()
}
