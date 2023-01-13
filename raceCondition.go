package main

import(
	"fmt"
	"time"
)

func goRoutine(count *int){

	for i:=0; i<10000; i++{
		*count ++
	}
}

func goRoutine2(count *int){

	for i:=0; i<10000; i++{
		*count --
	}
}

func main(){

	var count int
	go goRoutine(&count)
	go goRoutine2(&count)

	time.Sleep(10 * time.Millisecond)
	fmt.Println(count)
}