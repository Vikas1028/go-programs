package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	randomNumbers := map[string]int{}
	var countOfLessThan5Numbers, countOfLessThan10Numbers, countOfLessThan15Numbers, countOfLessThan20Numbers, countOfLessThan25Numbers int

	for randomNumbersCount := 0; randomNumbersCount < 1000; randomNumbersCount++ {
		randomNumber := rand.Intn(26)

		if randomNumber <= 5 {
			countOfLessThan5Numbers++
			randomNumbers[">0 && <=5"] = countOfLessThan5Numbers
		} else if randomNumber > 5 && randomNumber <= 10 {
			countOfLessThan10Numbers++
			randomNumbers[">5 && <=10"] = countOfLessThan10Numbers
		} else if randomNumber > 10 && randomNumber <= 15 {
			countOfLessThan15Numbers++
			randomNumbers[">10 && <=15"] = countOfLessThan15Numbers
		} else if randomNumber > 15 && randomNumber <= 20 {
			countOfLessThan20Numbers++
			randomNumbers[">15 && <=20"] = countOfLessThan20Numbers
		} else {
			countOfLessThan25Numbers++
			randomNumbers[">20 && <=25"] = countOfLessThan25Numbers
		}
	}

	fmt.Println("Count number of integers which are less than equal to 5: ", countOfLessThan5Numbers)
	fmt.Println("Count number of integers which are less than equal to 10: ", countOfLessThan10Numbers)
	fmt.Println("Count number of integers which are less than equal to 15: ", countOfLessThan15Numbers)
	fmt.Println("Count number of integers which are less than equal to 20: ", countOfLessThan20Numbers)
	fmt.Println("Count number of integers which are less than equal to 25: ", countOfLessThan25Numbers)

	// sliceOf5 := []int{}

	// for _, value := range randomNumbers {
	// 	if value <= 5 {
	// 		sliceOf5 = append(sliceOf5, value)
	// 	}
	// }

	fmt.Println(randomNumbers)
}
