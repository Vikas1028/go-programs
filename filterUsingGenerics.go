package main

import (
	"fmt"
	"math"
)

func filter[T any](slice []T, condition func(T) bool) []T {

	var filterSlice = []T{}

	for _, num := range slice {
		if condition(num) {
			filterSlice = append(filterSlice, num)
		}
	}

	return filterSlice

}

func noDecimal(num float64) bool {

	return math.Mod(num, 1) == 0
}

func divisibleByThree(num float64) bool {

	return math.Mod(num, 3) == 0
}

func startWithCharB(word string) bool {

	return word[0] == 66|98
}

func lengthThree(word string) bool {

	return len(word) == 3
}

func main() {

	sliceOfNum := []float64{5.12, 15.0, 3.39, 8.0, 9.0, 14.0, 6.63}

	fmt.Println("Slice Of Numbers: ", sliceOfNum)

	fmt.Println("Numbers which do not have any digit after decimal: ", filter(sliceOfNum, noDecimal))

	fmt.Println("Numbers which are divisible by 3: ", filter(sliceOfNum, divisibleByThree))

	fmt.Println("________________________________________________________________")

	sliceOfString := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	fmt.Println("Slice Of String: ", sliceOfString)

	fmt.Println("Words which start form character 'b': ", filter(sliceOfString, startWithCharB))

	fmt.Println("Words which have length 3: ", filter(sliceOfString, lengthThree))
}
