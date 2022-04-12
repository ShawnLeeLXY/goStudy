package concurrent

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func atomicDemo() {
	var num int64
	var wg sync.WaitGroup
	var lock sync.Mutex

	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go mutexAdd(&num, &lock, &wg)
	}
	wg.Wait()
	end := time.Now()
	fmt.Println("mutex用了", end.Sub(start))

	start = time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go atomicAdd(&num, &wg)
	}
	wg.Wait()
	end = time.Now()
	fmt.Println("atomic用了", end.Sub(start))
}

// 互斥锁
func mutexAdd(num *int64, lock *sync.Mutex, wg *sync.WaitGroup) {
	lock.Lock()
	*num += 1
	lock.Unlock()
	wg.Done()
}

// 原子操作
func atomicAdd(num *int64, wg *sync.WaitGroup) {
	atomic.AddInt64(num, 1)
	wg.Done()
}
