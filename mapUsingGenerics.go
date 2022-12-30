package main

import (
	"fmt"
	"strings"
)

func map1[T any](slice []T, condition func(T) T) []T {
	filteredSlice := []T{}

	for _, num := range slice {
		filteredSlice = append(filteredSlice, condition(num))
	}
	return filteredSlice
}

func main() {

	sliceOfNum := []float64{0.6, 0.3, 0.84, 0.04}

	fmt.Println("Slice of Numbers: ", sliceOfNum)

	fmt.Println("Convert each value as percentage: ", map1(sliceOfNum, func(num float64) float64 {
		return num * 100
	}))

	fmt.Println("compute half of each value: ", map1(sliceOfNum, func(num float64) float64 {
		return num / 2
	}))

	fmt.Println("______________________________________________________________")

	sliceOfString := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	fmt.Println("Slice Of String: ", sliceOfString)

	fmt.Println("First character of each word as capital: ", map1(sliceOfString, func(word string) string {
		return strings.Title(word)
	}))

	fmt.Println("Add colon ‘:’ at the end of word for the slice: ", map1(sliceOfString, func(word string) string {
		return fmt.Sprintf("%v:", word)
	}))

}
