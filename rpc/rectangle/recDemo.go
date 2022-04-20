package rectangle

import "sync"

var wg sync.WaitGroup

func RecDemo() {
	defer wg.Done()
	wg.Add(2)
	go server()
	go client()
	wg.Wait()
}
