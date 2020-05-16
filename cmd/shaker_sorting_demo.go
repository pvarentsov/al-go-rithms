package main

import (
	"fmt"
	"github.com/pvarentsov/al-go-rithms/internal/generator"
	"github.com/pvarentsov/al-go-rithms/internal/sorting"
	"time"
)

func main() {
	inputArray := generator.GenerateArray(10, true)

	fmt.Printf("Shaker sorting: %d\n\n", inputArray)
	time.Sleep(time.Millisecond * 2000)

	sorting.ShakerSortDemo(inputArray, -1, 200, 500)

	fmt.Print("Sorting finished!\n\n")
}
