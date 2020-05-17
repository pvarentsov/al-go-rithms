package sorting

func CombSort(inputArray []int, cmpClause int8) []int {
	sample := make([]int, len(inputArray))
	copy(sample, inputArray)

	if cmpClause == 0 {
		return sample
	}

	const factor = 1.247
	gap := len(inputArray) - 1

	for gap >= 1 {
		if cmpClause < 0 {
			for i := 0; i+gap < len(inputArray); i++ {
				if sample[i] > sample[i+gap] {
					sample[i], sample[i+gap] = sample[i+gap], sample[i]
				}
			}
		}
		if cmpClause > 0 {
			for i := 0; i+gap < len(inputArray); i++ {
				if sample[i] < sample[i+gap] {
					sample[i], sample[i+gap] = sample[i+gap], sample[i]
				}
			}
		}

		gap = int(float32(gap) / factor)
	}

	return sample
}
