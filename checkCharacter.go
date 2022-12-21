package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a String: ")
	scanner.Scan()
	nameString := scanner.Text()

	if (strings.Contains(nameString, "a") || strings.Contains(nameString, "A")) && (strings.Contains(nameString, "e") || strings.Contains(nameString, "E")) && (strings.Contains(nameString, "p") || strings.Contains(nameString, "P")) {
		fmt.Println("All Present")
	} else if strings.Contains(nameString, "a") || strings.Contains(nameString, "A") || strings.Contains(nameString, "e") || strings.Contains(nameString, "E") || strings.Contains(nameString, "p") || strings.Contains(nameString, "P") {
		fmt.Println("One or more - Present")
	} else {
		fmt.Println("None Present")
	}
}
