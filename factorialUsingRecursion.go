package main

import (
	"fmt"
)

/*
	    description: using recursion find the factorial of num
		param: num
		return (factorial number)int
*/
func computsFactorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * computsFactorial(num-1)
	}
}

func main() {
	var num int
	fmt.Print("Enter any number: ")
	fmt.Scanln(&num)
	fmt.Println(computsFactorial(num))
}
