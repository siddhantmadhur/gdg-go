package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"os"
)

func main() {
	val, err := russianRoulette(4)
	if err != nil {
		fmt.Printf("There was an error: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("Safe with val %d\n", val)
}

func russianRoulette(bullet int) (int, error) {
	if bullet >= 6 {
		return -1, errors.New("Bullet has to be between 0 and 5")
	}
	val := rand.Int() % 6
	fmt.Printf("Got: %d\n", val)
	if val == bullet {
		return val, errors.New("Lost!")
	}
	return val, nil
}
