package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
discription: Return count of vowels and consonants of user input string
param: sentance string
*/
func vowelsAndConsonants(sentence string) (int, int) {

	vowelsCount := 0
	consonantsCount := 0

	for _, char := range sentence {
		if char == 97 || char == 101 || char == 105 || char == 111 || char == 117 ||
			char == 65 || char == 69 || char == 73 || char == 79 || char == 85 {
			vowelsCount++
		} else if (char > 96 && char < 123) || (char > 64 && char < 91) {
			consonantsCount++
		}
	}

	return vowelsCount, consonantsCount
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a sentence: ")
	scanner.Scan()
	sentence := scanner.Text()

	vowelsCount, consonantsCount := vowelsAndConsonants(sentence)
	fmt.Printf("vowelsCount: %v \nconsonantsCount: %v", vowelsCount, consonantsCount)
}
