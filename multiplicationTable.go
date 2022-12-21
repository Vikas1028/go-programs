package main

import (
	"fmt"
)

func main() {
	var inputNum int
	fmt.Print("Enter any number between 2 to 25: ")
	fmt.Scanln(&inputNum)

	for i := 1; i <= 10; i++ {
		fmt.Println(inputNum, " X ", i, " = ", inputNum*i)
	}
}
