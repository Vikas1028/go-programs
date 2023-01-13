package main

import "fmt"

func evenChan(array []int, evenChannel chan int) {
	for _, num := range array {
		if num%2 == 0 {
			evenChannel <- num
		}
	}
	close(evenChannel)
}

func oddChan(array []int, oddChannel chan int) {
	for _, num := range array {
		if num%2 != 0 {
			oddChannel <- num
		}
	}
	close(oddChannel)
}

func main() {
	sliceOfNUmbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var evenNum, oddNum []int

	evenChannel := make(chan int)
	oddChannel := make(chan int)

	go evenChan(sliceOfNUmbers, evenChannel)
	go oddChan(sliceOfNUmbers, oddChannel)

	for num := range evenChannel {
		evenNum = append(evenNum, num)
	}

	for num := range oddChannel {
		oddNum = append(oddNum, num)
	}

	fmt.Println("Even numbers:", evenNum)
	fmt.Println("Odd numbers:", oddNum)
}
