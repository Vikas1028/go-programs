package main

import (
	"fmt"
	"sort"
	"sync"
)

var wg sync.WaitGroup

func sortSubarray1(arraySlice1 *[]float64) {

	sort.Float64s(*arraySlice1)
	wg.Done()
}

func sortSubarray2(arraySlice2 *[]float64) {

	sort.Float64s(*arraySlice2)
	wg.Done()
}

func sortSubarray3(arraySlice3 *[]float64) {

	sort.Float64s(*arraySlice3)
	wg.Done()
}

func sortSubarray4(arraySlice4 *[]float64) {

	sort.Float64s(*arraySlice4)
	wg.Done()
}

func main() {

	wg.Add(4)
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

	go sortSubarray1(&arraySlice1)
	go sortSubarray2(&arraySlice2)
	go sortSubarray3(&arraySlice3)
	go sortSubarray4(&arraySlice4)

	fmt.Println("Sorted Subarray: ", arraySlice1)
	fmt.Println("Sorted Subarray: ", arraySlice2)
	fmt.Println("Sorted Subarray: ", arraySlice3)
	fmt.Println("Sorted Subarray: ", arraySlice4)

	var sortedArray []float64
	sortedArray = append(sortedArray, arraySlice1...)
	sortedArray = append(sortedArray, arraySlice2...)
	sortedArray = append(sortedArray, arraySlice3...)
	sortedArray = append(sortedArray, arraySlice4...)

	sort.Float64s(sortedArray)
	fmt.Println("Sorted Array: ", sortedArray)
}
