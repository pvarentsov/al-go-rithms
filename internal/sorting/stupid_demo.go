package sorting

import (
	"github.com/gookit/color"
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/demo"

	"github.com/gosuri/uilive"
)

func StupidSortDemo(inputArray []int, cmpClause int8, indexDelay time.Duration, swapDelay time.Duration) []int {
	localArray := make([]int, len(inputArray))
	copy(localArray, inputArray)

	arrayDemo := demo.NewArrayDemo(localArray, uilive.New())

	if cmpClause == 0 {
		return localArray
	}

	index := 0

	arrayDemo.Render([]int{0}, []int{}, []int{}, indexDelay)

	for index < (len(localArray) - 1) {
		if cmpClause < 0 {
			if localArray[index] < localArray[index+1] {
				index++
				arrayDemo.Render([]int{index}, []int{}, []int{}, indexDelay)
			} else {
				arrayDemo.SetColor(color.Red)
				arrayDemo.Render([]int{}, []int{index, index + 1}, []int{}, swapDelay)

				localArray[index], localArray[index+1] = localArray[index+1], localArray[index]

				arrayDemo.SetColor(color.Green)
				arrayDemo.Render([]int{}, []int{index, index + 1}, []int{}, swapDelay)

				index = 0
				arrayDemo.Render([]int{index}, []int{}, []int{}, indexDelay)
			}
		}
		if cmpClause > 0 {
			if localArray[index] > localArray[index+1] {
				index++
				arrayDemo.Render([]int{index}, []int{}, []int{}, indexDelay)
			} else {
				arrayDemo.SetColor(color.Red)
				arrayDemo.Render([]int{}, []int{index, index + 1}, []int{}, swapDelay)

				localArray[index], localArray[index+1] = localArray[index+1], localArray[index]

				arrayDemo.SetColor(color.Green)
				arrayDemo.Render([]int{}, []int{index, index + 1}, []int{}, swapDelay)

				index = 0
				arrayDemo.Render([]int{index}, []int{}, []int{}, indexDelay)
			}
		}
	}

	arrayDemo.Close()

	return localArray
}
