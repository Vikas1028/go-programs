package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
description: count the vowels present in the user input string and then create map of vowels
param: inputString string
return: totalVowels and vowelsCountMap
*/
func vowelCount(inputString string) (int, map[string]int) {
	vowelsCountMap := map[string]int{}
	var countA, countE, countI, countO, countU int

	for _, char := range inputString {
		if char == 97 || char == 65 {
			countA++
			vowelsCountMap["vowel['a']"] = countA
		} else if char == 101 || char == 69 {
			countE++
			vowelsCountMap["vowel['e']"] = countE
		} else if char == 117 || char == 105 {
			countI++
			vowelsCountMap["vowel['i']"] = countI
		} else if char == 111 || char == 73 {
			countO++
			vowelsCountMap["vowel['o']"] = countO
		} else if char == 79 || char == 85 {
			countU++
			vowelsCountMap["vowel['u']"] = countU
		}

	}
	totalVowels := countA + countE + countI + countO + countU
	return totalVowels, vowelsCountMap
}

func main() {

	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter string: ")
	input.Scan()
	inputString := input.Text()

	vowels, vowelsMap := vowelCount(inputString)
	fmt.Println("Vowels: ", vowels)
	for vowel, count := range vowelsMap {
		fmt.Println(vowel, " : ", count)
	}

}
