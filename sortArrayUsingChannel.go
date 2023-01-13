package main

import (
	"fmt"
	"sort"
)

func sortSubarray1(arraySlice1 []float64, sortChan1 chan []float64) {

	sort.Float64s(arraySlice1)
	sortChan1 <- arraySlice1
}

func sortSubarray2(arraySlice2 []float64, sortChan2 chan []float64) {

	sort.Float64s(arraySlice2)
	sortChan2 <- arraySlice2
}

func sortSubarray3(arraySlice3 []float64, sortChan3 chan []float64) {

	sort.Float64s(arraySlice3)
	sortChan3 <- arraySlice3
}

func sortSubarray4(arraySlice4 []float64, sortChan4 chan []float64) {

	sort.Float64s(arraySlice4)
	sortChan4 <- arraySlice4
}

func main() {

	var arrayofNumbers []float64
	var size int
	var num float64

	fmt.Print("Enter size Of Array: ")
	fmt.Scanln(&size)
	fmt.Println("Enter array elements: ")
	for i := 0; i < size; i++ {
		fmt.Scanln(&num)
		arrayofNumbers = append(arrayofNumbers, num)
	}

	arraySlice1 := arrayofNumbers[0 : size/4]
	arraySlice2 := arrayofNumbers[size/4 : size/2]
	arraySlice3 := arrayofNumbers[size/2 : 3*size/4]
	arraySlice4 := arrayofNumbers[3*size/4 : size]

	sortChan1 := make(chan []float64)
	sortChan2 := make(chan []float64)
	sortChan3 := make(chan []float64)
	sortChan4 := make(chan []float64)

	go sortSubarray1(arraySlice1, sortChan1)
	go sortSubarray2(arraySlice2, sortChan2)
	go sortSubarray3(arraySlice3, sortChan3)
	go sortSubarray4(arraySlice4, sortChan4)

	sortSubarray1 := <-sortChan1
	sortSubarray2 := <-sortChan2
	sortSubarray3 := <-sortChan3
	sortSubarray4 := <-sortChan4

	fmt.Println("Sorted Subarray: ", sortSubarray1)
	fmt.Println("Sorted Subarray: ", sortSubarray2)
	fmt.Println("Sorted Subarray: ", sortSubarray3)
	fmt.Println("Sorted Subarray: ", sortSubarray4)

	var sortedArray []float64
	sortedArray = append(sortedArray, sortSubarray1...)
	sortedArray = append(sortedArray, sortSubarray2...)
	sortedArray = append(sortedArray, sortSubarray3...)
	sortedArray = append(sortedArray, sortSubarray4...)

	sort.Float64s(sortedArray)
	fmt.Println("Sorted Array: ", sortedArray)
}
