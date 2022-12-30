package main

import (
	"fmt"
	"strings"
)

func map1(sliceOfString []string, condition func(string) string) []string {

	var filterSliceOfString = []string{}

	for _, word := range sliceOfString {
		filterSliceOfString = append(filterSliceOfString, condition(word))
	}

	return filterSliceOfString
}

func main() {

	sliceOfString := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	fmt.Println("Slice Of String: ", sliceOfString)

	fmt.Println("First character of each word as capital: ", map1(sliceOfString, func(word string) string {
		return strings.Title(word)
	}))

	fmt.Println("Add colon ‘:’ at the end of word for the slice: ", map1(sliceOfString, func(word string) string {
		return fmt.Sprintf("%v:", word)
	}))
}
