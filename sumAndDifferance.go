package main

import (
	"fmt"
)

/*
discription: sum of two user input numbers
param: inputNum1, inputNum2 int
*/
func sum(inputNum1, inputNum2 int) int {
	return inputNum1 + inputNum2
}

/*
discription: diff of two user input numbers
param: inputNum1, inputNum2 int
*/
func diff(inputNum1, inputNum2 int) int {

	if inputNum1 > inputNum2 {
		return inputNum1 - inputNum2
	} else {
		return inputNum2 - inputNum1
	}
}

/*
discription: return two values
param: sum, diff int
*/
func Sum_n_Diff(sum, diff int) (int, int) {
	return sum, diff
}

func main() {

	var inputNum1, inputNum2 int
	fmt.Print("Enter any two Numbers: ")
	fmt.Scanln(&inputNum1, &inputNum2)

	fmt.Println("Sum: ", sum(inputNum1, inputNum2))
	fmt.Println("Difference: ", diff(inputNum1, inputNum2))

	sum := sum(inputNum1, inputNum2)
	diff := diff(inputNum1, inputNum2)
	sum2, diff2 := Sum_n_Diff(sum, diff)

	fmt.Printf("Sum and Diff: %v, %v", sum2, diff2)
}
