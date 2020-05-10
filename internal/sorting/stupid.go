package sorting

import (
	"github.com/gookit/color"
	"github.com/gosuri/uilive"
	"github.com/pvarentsov/al-go-rithms/internal/converter"
	"strings"
	"time"
)

// TODO: Code refactoring

func StupidSort(inputArray []int, cmpClause int8) []int {
	localArray := make([]int, len(inputArray))
	copy(localArray, inputArray)

	if cmpClause == 0 {
		return localArray
	}

	index := 0

	writer := uilive.New()
	writer.Start()

	strArray := generator.IntArrayToStrArray(localArray)
	strArray[0] = color.OpUnderscore.Sprintf("%s", strArray[0])

	color.Fprintf(writer, "%s\n\n", strings.Join(strArray, " "))
	time.Sleep(time.Millisecond * 1000)

	for index < (len(localArray) - 1) {
		if cmpClause < 0 {
			if localArray[index] < localArray[index+1] {
				index++

				strArray := generator.IntArrayToStrArray(localArray)
				strArray[index] = color.OpUnderscore.Sprintf("%s", strArray[index])

				color.Fprintf(writer, "%s\n\n", strings.Join(strArray, " "))
				time.Sleep(time.Millisecond * 1000)

			} else {
				localArray[index], localArray[index+1] = localArray[index+1], localArray[index]

				strArray := generator.IntArrayToStrArray(localArray)
				strArray[index] = color.Red.Sprintf("%s", strArray[index])
				strArray[index+1] = color.Red.Sprintf("%s", strArray[index+1])

				color.Fprintf(writer, "%s\n\n", strings.Join(strArray, " "))
				time.Sleep(time.Millisecond * 1000)

				index = 0

				strArrayAfterReset := generator.IntArrayToStrArray(localArray)
				strArrayAfterReset[index] = color.OpUnderscore.Sprintf("%s", strArrayAfterReset[index])

				color.Fprintf(writer, "%s\n\n", strings.Join(strArrayAfterReset, " "))
				time.Sleep(time.Millisecond * 1000)
			}
		}
		if cmpClause > 0 {
			if localArray[index] > localArray[index+1] {
				index++
			} else {
				localArray[index], localArray[index+1] = localArray[index+1], localArray[index]
				index = 0
			}
		}
	}

	writer.Stop()

	return localArray
}
