package main

import (
	"fmt"
	"math"
)

func filter(sliceOfNum []float64, condition func(float64) bool) []float64 {

	var filterSliceOfNum = []float64{}

	for _, num := range sliceOfNum {
		if condition(num) {
			filterSliceOfNum = append(filterSliceOfNum, num)
		}
	}

	return filterSliceOfNum

}

func noDecimal(num float64) bool {

	return math.Mod(num, 1) == 0
}

func divisibleByThree(num float64) bool {

	return math.Mod(num, 3) == 0
}

func main() {

	sliceOfNum := []float64{5.12, 15.0, 3.39, 8.0, 9.0, 14.0, 6.63}

	fmt.Println("Slice Of Numbers: ", sliceOfNum)

	fmt.Println("Numbers which do not have any digit after decimal: ", filter(sliceOfNum, noDecimal))

	fmt.Println("Numbers which are divisible by 3: ", filter(sliceOfNum, divisibleByThree))
}
