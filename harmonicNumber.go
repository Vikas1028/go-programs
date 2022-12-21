package main

import (
	"fmt"
)

/*
	    desc: to find the harmonic number of user entered number
		param: number
*/
func findHarmonicNumber(number int) {

	var harmonicNumber float64
	for i := 1; i < number; i++ {
		harmonicNumber += 1 / float64(i)
	}
	fmt.Println("Harmonic Number: ", harmonicNumber)
}

func main() {

	var number int
	fmt.Print("Enter any Number: ")
	fmt.Scanln(&number)
	(findHarmonicNumber(number))
}
