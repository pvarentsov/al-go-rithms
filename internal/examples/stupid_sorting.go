package examples

import (
	"fmt"
	"github.com/pvarentsov/al-go-rithms/internal/generator"
	"github.com/pvarentsov/al-go-rithms/internal/sorting"
	"time"
)

func ShowStupidSorting() {
	inputArray := generator.GenerateArray(10, true)

	fmt.Printf("Stupid sorting: %d\n\n", inputArray)
	time.Sleep(time.Millisecond * 2000)

	sorting.StupidSort(inputArray, -1)

	fmt.Println("\nSorting finished!")
}
