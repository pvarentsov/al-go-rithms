package sorting

func BubbleSort(inputArray []int, cmpClause int8) []int {
	sample := make([]int, len(inputArray))
	copy(sample, inputArray)

	if cmpClause == 0 {
		return sample
	}

	if cmpClause < 0 {
		for i := 0; i < (len(sample) - 1); i++ {
			for j := 0; j < (len(sample) - i - 1); j++ {
				if sample[j] > sample[j+1] {
					sample[j], sample[j+1] = sample[j+1], sample[j]
				}
			}
		}
	}
	if cmpClause > 0 {
		for i := 0; i < (len(sample) - 1); i++ {
			for j := 0; j < (len(sample) - i - 1); j++ {
				if sample[j] < sample[j+1] {
					sample[j], sample[j+1] = sample[j+1], sample[j]
				}
			}
		}
	}

	return sample
}
