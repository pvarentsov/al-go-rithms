package sorting

func StupidSort(inputArray []int, cmpClause int8) []int {
	sample := make([]int, len(inputArray))
	copy(sample, inputArray)

	if cmpClause == 0 {
		return sample
	}

	index := 0

	for index < (len(sample) - 1) {
		if cmpClause < 0 {
			if sample[index] < sample[index+1] {
				index++
			} else {
				sample[index], sample[index+1] = sample[index+1], sample[index]
				index = 0
			}
		}
		if cmpClause > 0 {
			if sample[index] > sample[index+1] {
				index++
			} else {
				sample[index], sample[index+1] = sample[index+1], sample[index]
				index = 0
			}
		}
	}

	return sample
}
