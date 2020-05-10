package sorting

func StupidSort(inputArray []int, cmpClause int8) []int {
	localArray := make([]int, len(inputArray))
	copy(localArray, inputArray)

	if cmpClause == 0 {
		return localArray
	}

	index := 0

	for index < (len(localArray) - 1) {
		if cmpClause < 0 {
			if localArray[index] < localArray[index+1] {
				index++
			} else {
				localArray[index], localArray[index+1] = localArray[index+1], localArray[index]
				index = 0
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

	return localArray
}
