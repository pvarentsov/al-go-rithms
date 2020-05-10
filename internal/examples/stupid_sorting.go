package examples

import (
	"fmt"
	"github.com/pvarentsov/al-go-rithms/internal/generator"
	"github.com/pvarentsov/al-go-rithms/internal/sorting"
)

func ShowStupidSorting() {
	inputArray := generator.GenerateArray(20, true)
	sortedArray := sorting.StupidSort(inputArray, -1)

	fmt.Println("Stupid sorting:")
	fmt.Println()

	fmt.Printf("Input Array : %d", inputArray)
	fmt.Println()
	fmt.Printf("Sorted Array: %d", sortedArray)
	fmt.Println()
}
