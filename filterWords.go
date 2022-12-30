package main

import (
	"fmt"
)

func filter(sliceOfString []string, condition func(word string) bool) []string {

	var filterSliceOfString = []string{}

	for _, word := range sliceOfString {
		if condition(word) {
			filterSliceOfString = append(filterSliceOfString, word)
		}
	}

	return filterSliceOfString

}

func startWithCharB(word string) bool {

	return word[0] == 66|98
}

func lengthThree(word string) bool {

	return len(word) == 3
}

func main() {

	sliceOfString := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	fmt.Println("Slice Of String: ", sliceOfString)

	fmt.Println("Words which start form character 'b': ", filter(sliceOfString, startWithCharB))

	fmt.Println("Words which have length 3: ", filter(sliceOfString, lengthThree))
}
