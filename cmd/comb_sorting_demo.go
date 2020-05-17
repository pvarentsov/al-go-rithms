package main

import (
	"fmt"
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/generator"
	"github.com/pvarentsov/al-go-rithms/internal/sorting"
)

func main() {
	inputArray := generator.GenerateArray(10, true)

	fmt.Printf("Comb sorting: %d\n\n", inputArray)
	time.Sleep(time.Millisecond * 2000)

	sorting.CombSortDemo(inputArray, -1, 200, 500)

	fmt.Print("Sorting finished!\n\n")
}
