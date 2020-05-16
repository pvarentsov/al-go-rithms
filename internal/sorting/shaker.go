package sorting

func ShakerSort(inputArray []int, cmpClause int8) []int {
	sample := make([]int, len(inputArray))
	copy(sample, inputArray)

	if cmpClause == 0 {
		return sample
	}

	leftIndex := 0
	rightIndex := len(sample) - 1

	for leftIndex <= rightIndex {
		if cmpClause < 0 {
			for i := leftIndex; i < rightIndex; i++ {
				if sample[i] > sample[i+1] {
					sample[i], sample[i+1] = sample[i+1], sample[i]
				}
			}

			rightIndex -= 1

			for i := rightIndex; i > leftIndex; i-- {
				if sample[i-1] > sample[i] {
					sample[i], sample[i-1] = sample[i-1], sample[i]
				}
			}

			leftIndex += 1
		}
		if cmpClause > 0 {
			for i := leftIndex; i < rightIndex; i++ {
				if sample[i] < sample[i+1] {
					sample[i], sample[i+1] = sample[i+1], sample[i]
				}
			}

			rightIndex -= 1

			for i := rightIndex; i > leftIndex; i-- {
				if sample[i-1] < sample[i] {
					sample[i], sample[i-1] = sample[i-1], sample[i]
				}
			}

			leftIndex += 1
		}
	}

	return sample
}
