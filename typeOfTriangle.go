package main

import (
	"fmt"
)

func main() {
	var inputNum, inputNum2, inputNum3 int
	fmt.Print("Enter length of 3 sids of Triangle: ")
	fmt.Scanln(&inputNum, &inputNum2, &inputNum3)

	if inputNum == inputNum2 && inputNum == inputNum3 {
		fmt.Println("Equilateral Triangle")
	} else if inputNum == inputNum2 || inputNum == inputNum3 || inputNum2 == inputNum3 {
		fmt.Println("Isosceles Triangle")
	} else {
		fmt.Println("Scalene Triangle")
	}

}
