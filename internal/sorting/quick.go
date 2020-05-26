package sorting

// TODO: How quick sort can be visualized?

func QuickSort(sample []int, cmpClause int8) []int {
	if cmpClause == 0 {
		return sample
	}
	if len(sample) < 2 {
		return sample
	}

	pivotIndex := len(sample) >> 1
	pivot := sample[pivotIndex]

	left, right := make([]int, 0), make([]int, 0)

	for _, v := range sample {
		if cmpClause < 0 {
			if v < pivot {
				left = append(left, v)
			}
			if v > pivot {
				right = append(right, v)
			}
		}
		if cmpClause > 0 {
			if v < pivot {
				right = append(right, v)
			}
			if v > pivot {
				left = append(left, v)
			}
		}
	}

	left, right = QuickSort(left, cmpClause), QuickSort(right, cmpClause)

	sorted := append(left, pivot)
	sorted = append(sorted, right...)

	return sorted
}
