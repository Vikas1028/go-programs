package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bufferedChan(ch chan int) {

	for i := 0; i < 15; i++ {
		ch <- rand.Intn(100)
		time.Sleep(time.Second)
	}
	close(ch)
}

func main() {

	ch := make(chan int, 3)
	go bufferedChan(ch)

	for num := range ch {
		fmt.Println(num)
	}
}
