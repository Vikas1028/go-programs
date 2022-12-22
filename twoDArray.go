package main

import (
	"fmt"
)

func main() {

	var _2DArray = [5][2]int{}
	_2DArray[0][0] = 10
	_2DArray[1][0] = 12
	_2DArray[2][0] = 15
	_2DArray[3][0] = 19
	_2DArray[4][0] = 24

	for i := 0; i < len(_2DArray); i++ {
		_2DArray[i][1] = _2DArray[i][0] * 2
	}
	fmt.Println(_2DArray)

}
