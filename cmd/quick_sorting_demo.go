package main

import (
	"fmt"

	"github.com/pvarentsov/al-go-rithms/internal/generator"
	"github.com/pvarentsov/al-go-rithms/internal/sorting"
)

func main() {
	inputArray := generator.GenerateArray(10, true)

	fmt.Printf("Quick sorting: %d\n\n", inputArray)

	inputArray = sorting.QuickSort(inputArray, -1)

	fmt.Printf("%d\n\n", inputArray)
}
