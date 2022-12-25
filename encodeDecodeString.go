package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
description: encode string of user input
param: inputString string
return: slice of []rune
*/
func encode(inputString string) []rune {

	encodeSlice := make([]rune, len(inputString))
	for indexOfChar, char := range inputString {
		encodeSlice[indexOfChar] = char + 1
	}
	return encodeSlice
}

/*
description: decode the encode Slice
param: encode []rune
return: decode slice []rune
*/
func decode(encode []rune) []rune {

	decodeSlice := make([]rune, len(encode))
	for indexOfChar, char := range encode {
		decodeSlice[indexOfChar] = char - 1
	}
	return decodeSlice
}

func main() {
	fmt.Print("Enter string to encode and decode: ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	inputString := input.Text()

	encode := encode(inputString)
	decode := decode(encode)
	fmt.Printf("Encode String: %v \nDecode String: %v \nMatch inputString and DecodeString: %v", string(encode), string(decode), inputString == string(decode) /*strings.Compare(inputString, string(decode)) == 0*/)
}
