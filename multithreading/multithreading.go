package main

import (
	"fmt"
	"time"
)

func main() {
	var sum int = 0

	for range 5 {
		add(&sum)
	}

	fmt.Printf("Sum: %d\n", sum)
}

func add(sum *int) {
	time.Sleep(time.Second)
	*sum += 1
}
