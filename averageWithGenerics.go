package main

import "fmt"

func average[T int | float64](num ...T) string {

	var total T
	for _, num := range num {
		total += num
	}
	average := total / T(len(num))

	return fmt.Sprintf("Average of %v is: %v", num, average)
}

func main() {

	fmt.Println(average(6, 4, 7, 8))
	fmt.Println(average(6.3, 4.3, 7.5, 8.9, 10.1, 20.3))
	fmt.Println(average(66.76, 42.43, 7.32))
	fmt.Println(average(6, 4, 7, 8, 25))
	fmt.Println(average(6, 4))

}
