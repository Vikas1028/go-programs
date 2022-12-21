package main

import (
	"fmt"
)

func computsFactorialOfNumber(num int) int {
	sum := 1
	for num > 0 {
		sum = sum * num
		num--
	}
	return sum
}

func main() {
	var num int
	fmt.Print("Enter any number: ")
	fmt.Scanln(&num)
	fmt.Println(computsFactorialOfNumber(num))
}
