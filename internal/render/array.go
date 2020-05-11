package render

import (
	"io"
	"strings"
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/converter"

	"github.com/gookit/color"
)

func MarkArrayIndex(array []int, index int, delayInMs time.Duration, writer io.Writer, boldIndexes []int) {
	strArray := generator.IntArrayToStrArray(array)
	strArray[index] = color.OpUnderscore.Sprintf("%s", strArray[index])

	for _, v := range boldIndexes {
		strArray[v] = color.Bold.Sprintf("%s", strArray[v])
	}

	color.Fprintf(writer, "%s\n\n", strings.Join(strArray, " "))
	time.Sleep(time.Millisecond * delayInMs)
}

func SwapArrayElements(array []int, leftIndex int, rightIndex int, delayInMs time.Duration, switchColors bool, writer io.Writer, boldIndexes []int) {
	leftIndexColor := color.Red
	rightIndexColor := color.Red

	if switchColors {
		leftIndexColor, rightIndexColor = color.Green, color.Green
	}

	strArray := generator.IntArrayToStrArray(array)
	strArray[leftIndex] = leftIndexColor.Sprintf("%s", strArray[leftIndex])
	strArray[rightIndex] = rightIndexColor.Sprintf("%s", strArray[rightIndex])

	for _, v := range boldIndexes {
		strArray[v] = color.Bold.Sprintf("%s", strArray[v])
	}

	color.Fprintf(writer, "%s\n\n", strings.Join(strArray, " "))
	time.Sleep(time.Millisecond * delayInMs)
}
