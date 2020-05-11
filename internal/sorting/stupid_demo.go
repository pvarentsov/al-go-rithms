package sorting

import (
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/render"

	"github.com/gosuri/uilive"
)

func StupidSortDemo(inputArray []int, cmpClause int8, indexDelay time.Duration, swapDelay time.Duration) []int {
	localArray := make([]int, len(inputArray))
	copy(localArray, inputArray)

	if cmpClause == 0 {
		return localArray
	}

	index := 0

	writer := uilive.New()
	writer.Start()

	render.MarkArrayIndex(localArray, 0, indexDelay, writer, make([]int, 0))

	for index < (len(localArray) - 1) {
		if cmpClause < 0 {
			if localArray[index] < localArray[index+1] {
				index++
				render.MarkArrayIndex(localArray, index, indexDelay, writer, make([]int, 0))
			} else {
				render.SwapArrayElements(localArray, index, index+1, swapDelay, false, writer, make([]int, 0))
				localArray[index], localArray[index+1] = localArray[index+1], localArray[index]
				render.SwapArrayElements(localArray, index, index+1, swapDelay, true, writer, make([]int, 0))

				index = 0
				render.MarkArrayIndex(localArray, index, indexDelay, writer, make([]int, 0))
			}
		}
		if cmpClause > 0 {
			if localArray[index] > localArray[index+1] {
				index++
				render.MarkArrayIndex(localArray, index, indexDelay, writer, make([]int, 0))
			} else {
				render.SwapArrayElements(localArray, index, index+1, swapDelay, false, writer, make([]int, 0))
				localArray[index], localArray[index+1] = localArray[index+1], localArray[index]
				render.SwapArrayElements(localArray, index, index+1, swapDelay, true, writer, make([]int, 0))

				index = 0
				render.MarkArrayIndex(localArray, index, indexDelay, writer, make([]int, 0))
			}
		}
	}

	writer.Stop()

	return localArray
}
