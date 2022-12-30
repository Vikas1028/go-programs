package main

import "fmt"

func average(num ...int) string {

	total := 0
	for _, num := range num {
		total += num
	}
	average := total / len(num)

	return fmt.Sprintf("Average of %v is: %v", num, average)
}

func main() {

	fmt.Println(average(6, 4, 7, 8))
	fmt.Println(average(6, 4, 7, 8, 10, 20))
	fmt.Println(average(6, 4, 7))
	fmt.Println(average(6, 4, 7, 8, 25))
	fmt.Println(average(6, 4))

}
