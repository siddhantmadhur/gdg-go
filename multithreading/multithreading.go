package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var sum int = 0

	var wg sync.WaitGroup
	var m sync.Mutex
	for range 5 {
		wg.Add(1)
		go add(&sum, &wg, &m)
	}
	wg.Wait()

	fmt.Printf("Sum: %d\n", sum)
}

func add(sum *int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	time.Sleep(time.Second)
	m.Lock()
	*sum += 1
	m.Unlock()
}
