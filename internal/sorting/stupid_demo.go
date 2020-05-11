package sorting

import (
	"io"
	"strings"
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/converter"

	"github.com/gookit/color"
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

	renderMarkingArrayIndex(localArray, 0, indexDelay, writer)

	for index < (len(localArray) - 1) {
		if cmpClause < 0 {
			if localArray[index] < localArray[index+1] {
				index++
				renderMarkingArrayIndex(localArray, index, indexDelay, writer)
			} else {
				renderSwappingArrayElements(localArray, index, index+1, swapDelay, false, writer)
				localArray[index], localArray[index+1] = localArray[index+1], localArray[index]
				renderSwappingArrayElements(localArray, index, index+1, swapDelay, true, writer)

				index = 0
				renderMarkingArrayIndex(localArray, index, indexDelay, writer)
			}
		}
		if cmpClause > 0 {
			if localArray[index] > localArray[index+1] {
				index++
				renderMarkingArrayIndex(localArray, index, indexDelay, writer)
			} else {
				renderSwappingArrayElements(localArray, index, index+1, swapDelay, false, writer)
				localArray[index], localArray[index+1] = localArray[index+1], localArray[index]
				renderSwappingArrayElements(localArray, index, index+1, swapDelay, true, writer)

				index = 0
				renderMarkingArrayIndex(localArray, index, indexDelay, writer)
			}
		}
	}

	writer.Stop()

	return localArray
}

func renderMarkingArrayIndex(array []int, index int, delayInMs time.Duration, writer io.Writer) {
	strArray := generator.IntArrayToStrArray(array)
	strArray[index] = color.OpUnderscore.Sprintf("%s", strArray[index])

	color.Fprintf(writer, "%s\n\n", strings.Join(strArray, " "))
	time.Sleep(time.Millisecond * delayInMs)
}

func renderSwappingArrayElements(array []int, leftIndex int, rightIndex int, delayInMs time.Duration, switchColors bool, writer io.Writer) {
	leftIndexColor := color.Red
	rightIndexColor := color.Red

	if switchColors {
		leftIndexColor, rightIndexColor = color.Green, color.Green
	}

	strArray := generator.IntArrayToStrArray(array)
	strArray[leftIndex] = leftIndexColor.Sprintf("%s", strArray[leftIndex])
	strArray[rightIndex] = rightIndexColor.Sprintf("%s", strArray[rightIndex])

	color.Fprintf(writer, "%s\n\n", strings.Join(strArray, " "))
	time.Sleep(time.Millisecond * delayInMs)
}
