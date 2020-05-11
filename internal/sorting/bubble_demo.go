package sorting

import (
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/render"

	"github.com/gosuri/uilive"
)

func BubbleSortDemo(inputArray []int, cmpClause int8, indexDelay time.Duration, swapDelay time.Duration) []int {
	localArray := make([]int, len(inputArray))
	copy(localArray, inputArray)

	writer := uilive.New()
	writer.Start()

	if cmpClause == 0 {
		return localArray
	}

	alreadySortedIndexes := make([]int, 0)

	if cmpClause < 0 {
		for i := 0; i < (len(localArray) - 1); i++ {
			for j := 0; j < (len(localArray) - i - 1); j++ {
				render.MarkArrayIndex(localArray, j, indexDelay, writer, alreadySortedIndexes)

				if localArray[j] > localArray[j+1] {
					render.SwapArrayElements(localArray, j, j+1, swapDelay, false, writer, alreadySortedIndexes)
					localArray[j], localArray[j+1] = localArray[j+1], localArray[j]
					render.SwapArrayElements(localArray, j, j+1, swapDelay, true, writer, alreadySortedIndexes)
				}
				if j+1 == (len(localArray) - i - 1) {
					alreadySortedIndexes = append(alreadySortedIndexes, j+1)
				}
				if i+1 == (len(localArray) - 1) {
					render.MarkArrayIndex(localArray, j, indexDelay, writer, alreadySortedIndexes)
				}
			}
		}
	}
	if cmpClause > 0 {
		for i := 0; i < (len(localArray) - 1); i++ {
			for j := 0; j < (len(localArray) - i - 1); j++ {
				render.MarkArrayIndex(localArray, j, indexDelay, writer, alreadySortedIndexes)

				if localArray[j] < localArray[j+1] {
					render.SwapArrayElements(localArray, j, j+1, swapDelay, false, writer, alreadySortedIndexes)
					localArray[j], localArray[j+1] = localArray[j+1], localArray[j]
					render.SwapArrayElements(localArray, j, j+1, swapDelay, true, writer, alreadySortedIndexes)
				}
				if j+1 == (len(localArray) - i - 1) {
					alreadySortedIndexes = append(alreadySortedIndexes, j+1)
				}
				if i+1 == (len(localArray) - 1) {
					render.MarkArrayIndex(localArray, j, indexDelay, writer, alreadySortedIndexes)
				}
			}
		}
	}

	writer.Stop()

	return localArray
}
