package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
description: compute Return max and second Max number of slice
param: stringNumbers []string
*/
func maxNumbers(stringOfNumbers []string) (float64, float64) {
	maxNum := 1.0
	secondMaxNum := 0.0

	for i := 0; i < len(stringOfNumbers); i++ {
		num, _ := strconv.ParseFloat(stringOfNumbers[i], 64)
		if num > secondMaxNum {
			secondMaxNum = num
			if secondMaxNum >= maxNum {
				temp := maxNum
				maxNum = secondMaxNum
				secondMaxNum = temp
			}
		}
	}
	return maxNum, secondMaxNum
}

/*
description: compute Return min and second min number of slice
param: stringNumbers []string
*/
func minNumbers(stringOfNumbers []string) {

	var firstMinNum float64 = 1.7976931348623157e+308
	var secondMinNum float64 = 1.7976931348623157e+308
	for i := 0; i < len(stringOfNumbers); i++ {
		num, _ := strconv.ParseFloat(stringOfNumbers[i], 64)
		if num < secondMinNum {
			secondMinNum = num
			if secondMinNum <= firstMinNum {
				temp := firstMinNum
				firstMinNum = secondMinNum
				secondMinNum = temp
			}
		}
	}
	fmt.Println("minNum: ", firstMinNum)
	fmt.Println("secondMinNum: ", secondMinNum)
}

func main() {

	var commaSeparatedNumbers string
	fmt.Print("Enter sequence of comma-separated numbers: ")
	fmt.Scanln(&commaSeparatedNumbers)

	stringOfNumbers := strings.Split(commaSeparatedNumbers, ",")
	fmt.Println("Slice of Numbers: ", stringOfNumbers)

	maxNum, secMaxNum := maxNumbers(stringOfNumbers)
	fmt.Printf("Max Number: %v \nSecond Max Number: %v", maxNum, secMaxNum)

	minNumbers(stringOfNumbers)
}
