package main

import (
	"fmt"
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/generator"
	"github.com/pvarentsov/al-go-rithms/internal/sorting"

	"github.com/gosuri/uilive"
	"github.com/pvarentsov/al-go-rithms/internal/demo"
)

func main() {
	sortedArray := generator.GenerateArray(10, true)

	fmt.Printf("Quick sorting: %d\n\n", sortedArray)
	time.Sleep(time.Millisecond * 2000)

	sortedArray = sorting.QuickSort(sortedArray, -1)

	arrayDemo := demo.NewArrayDemo(sortedArray, uilive.New())
	arrayDemo.Render([]int{len(sortedArray) - 1}, []int{}, []int{}, 200)
	arrayDemo.Close()

	fmt.Print("Wow! Very quick!\n\n")
}
