package sorting

func BubbleSort(inputArray []int, cmpClause int8) []int {
	localArray := make([]int, len(inputArray))
	copy(localArray, inputArray)

	if cmpClause == 0 {
		return localArray
	}

	if cmpClause < 0 {
		for i := 0; i < (len(localArray) - 1); i++ {
			for j := 0; j < (len(localArray) - i - 1); j++ {
				if localArray[j] > localArray[j+1] {
					localArray[j], localArray[j+1] = localArray[j+1], localArray[j]
				}
			}
		}
	}
	if cmpClause > 0 {
		for i := 0; i < (len(localArray) - 1); i++ {
			for j := 0; j < (len(localArray) - i - 1); j++ {
				if localArray[j] < localArray[j+1] {
					localArray[j], localArray[j+1] = localArray[j+1], localArray[j]
				}
			}
		}
	}

	return localArray
}
