package main

import (
	"fmt"
	"testing"
)

func TestRoulette(t *testing.T) {
	_, err := russianRoulette(16)
	if err == nil {
		fmt.Printf("roulette doesnt error on out of bounds!\n")
		t.FailNow()
	}
	fmt.Printf("Passed roulette!\n")
}
