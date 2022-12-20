package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Simulate flipping of a Coin")

	for i := 0; i < 10; i++ {
		toss := (rand.Intn(2))
		if toss == 0 {
			fmt.Println("Tail")
		} else {
			fmt.Println("Head")
		}
	}

}
