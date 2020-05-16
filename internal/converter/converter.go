package converter

import (
	"fmt"
)

func IntArrayToStrArray(intArray []int) []string {
	stringArray := make([]string, len(intArray))

	for i := 0; i < len(stringArray); i++ {
		stringArray[i] = fmt.Sprintf("%d", intArray[i])
	}

	return stringArray
}
