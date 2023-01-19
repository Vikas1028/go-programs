package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var chopsticks [5]sync.Mutex

func philosopher(id int) {

	for i := 0; i < 3; i++ {
		chopsticks[id-1].Lock()
		chopsticks[(id)%5].Lock()
		fmt.Printf("philosopher %d starting to eat\n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("philosopher %d finishing eating\n", id)
		chopsticks[id-1].Unlock()
		chopsticks[(id)%5].Unlock()
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

func main() {

	wg.Add(5)
	for i := 1; i < 6; i++ {
		go philosopher(i)
	}
	wg.Wait()
}
