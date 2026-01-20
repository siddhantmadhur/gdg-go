package main

import (
	"fmt"
	"sync"
	"time"
)

var store = make(map[string]int)

var SLOW = true

func main() {
	go put("a", 120) // runs on a new thread
	put("b", 250)
	put("c", 140)

	var a, b, c int

	if SLOW {
		get("a", &a)
		get("b", &b)
		get("c", &c)
	} else {
		var wg sync.WaitGroup
		wg.Go(func() { get("a", &a) })
		wg.Go(func() { get("b", &b) })
		wg.Go(func() { get("c", &c) })
		wg.Wait()
	}

	sum := a + b + c

	fmt.Printf("Sum: %d\n", sum)
}

func put(key string, val int) {
	fmt.Printf("Updated (%s:%d)\n", key, val)
	store[key] = val
}

func get(key string, val *int) {
	time.Sleep(time.Second * 2)
	*val = store[key]
}
