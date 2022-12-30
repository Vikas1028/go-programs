package main

import "fmt"

func map1(sliceOfNum []float64, condition func(float64) float64) []float64 {
	filteredSliceOfNum := []float64{}

	for _, num := range sliceOfNum {
		filteredSliceOfNum = append(filteredSliceOfNum, condition(num))
	}
	return filteredSliceOfNum
}

func main() {

	sliceOfNum := []float64{0.6, 0.3, 0.84, 0.04}

	fmt.Println("Slice of Numbers: ", sliceOfNum)

	fmt.Println("Convert each value as percentage: ", map1(sliceOfNum, func(num float64) float64 {
		return num * 100
	}))

	fmt.Println("compute half of each value: ", map1(sliceOfNum, func(num float64) float64 {
		return num / 2
	}))

}
