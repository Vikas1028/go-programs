package main

import (
	"fmt"
)

/*
Description: Generate Odd number slice
param: arrSlice []int
*/
func isOdd(arrSlice []int) {
	oddSlice := [5]int{}
	j := 0
	for i := 0; i < len(arrSlice); i++ {
		if arrSlice[i]%2 != 0 {
			oddSlice[j] = arrSlice[i]
			j++
		}
	}
	fmt.Println("oddSlice: ", oddSlice)
}

/*
Description: Generate even number slice
param: arrSlice []int
*/
func isEven(arrSlice []int) {
	evenSlice := [5]int{}
	j := 0
	for i := 0; i < len(arrSlice); i++ {
		if arrSlice[i]%2 == 0 {
			evenSlice[j] = arrSlice[i]
			j++
		}
	}
	fmt.Println("evenSlice: ", evenSlice)
}

func main() {

	arr := [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	arrSlice := arr[4:14]

	isOdd(arrSlice)
	isEven(arrSlice)
}
