package sorting

import (
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/demo"

	"github.com/gookit/color"
	"github.com/gosuri/uilive"
)

func ShakerSortDemo(inputArray []int, cmpClause int8, indexDelay time.Duration, swapDelay time.Duration) []int {
	sample := make([]int, len(inputArray))
	copy(sample, inputArray)

	if cmpClause == 0 {
		return sample
	}

	arrayDemo := demo.NewArrayDemo(sample, uilive.New())

	alreadySortedIndexes := make([]int, 0)

	leftIndex := 0
	rightIndex := len(sample) - 1

	for leftIndex <= rightIndex {
		if cmpClause < 0 {
			for i := leftIndex; i < rightIndex; i++ {
				arrayDemo.Render([]int{i}, []int{}, alreadySortedIndexes, indexDelay)

				if sample[i] > sample[i+1] {
					arrayDemo.SetColor(color.Red)
					arrayDemo.Render([]int{}, []int{i, i + 1}, alreadySortedIndexes, swapDelay)

					sample[i], sample[i+1] = sample[i+1], sample[i]

					arrayDemo.SetColor(color.Green)
					arrayDemo.Render([]int{}, []int{i, i + 1}, alreadySortedIndexes, swapDelay)
				}
				if i+1 == rightIndex {
					alreadySortedIndexes = append(alreadySortedIndexes, rightIndex)
					arrayDemo.Render([]int{i}, []int{}, alreadySortedIndexes, indexDelay)
				}
			}

			rightIndex -= 1

			for i := rightIndex; i > leftIndex; i-- {
				arrayDemo.Render([]int{i}, []int{}, alreadySortedIndexes, indexDelay)

				if sample[i-1] > sample[i] {
					arrayDemo.SetColor(color.Red)
					arrayDemo.Render([]int{}, []int{i, i - 1}, alreadySortedIndexes, swapDelay)

					sample[i], sample[i-1] = sample[i-1], sample[i]

					arrayDemo.SetColor(color.Green)
					arrayDemo.Render([]int{}, []int{i, i - 1}, alreadySortedIndexes, swapDelay)
				}
				if i-1 == leftIndex {
					alreadySortedIndexes = append(alreadySortedIndexes, leftIndex)
					arrayDemo.Render([]int{i}, []int{}, alreadySortedIndexes, indexDelay)
				}
			}

			leftIndex += 1
		}
		if cmpClause > 0 {
			for i := leftIndex; i < rightIndex; i++ {
				arrayDemo.Render([]int{i}, []int{}, alreadySortedIndexes, indexDelay)

				if sample[i] < sample[i+1] {
					arrayDemo.SetColor(color.Red)
					arrayDemo.Render([]int{}, []int{i, i + 1}, alreadySortedIndexes, swapDelay)

					sample[i], sample[i+1] = sample[i+1], sample[i]

					arrayDemo.SetColor(color.Green)
					arrayDemo.Render([]int{}, []int{i, i + 1}, alreadySortedIndexes, swapDelay)
				}
				if i+1 == rightIndex {
					alreadySortedIndexes = append(alreadySortedIndexes, rightIndex)
					arrayDemo.Render([]int{i}, []int{}, alreadySortedIndexes, indexDelay)
				}
			}

			rightIndex -= 1

			for i := rightIndex; i > leftIndex; i-- {
				arrayDemo.Render([]int{i}, []int{}, alreadySortedIndexes, indexDelay)

				if sample[i-1] < sample[i] {
					arrayDemo.SetColor(color.Red)
					arrayDemo.Render([]int{}, []int{i, i - 1}, alreadySortedIndexes, swapDelay)

					sample[i], sample[i-1] = sample[i-1], sample[i]

					arrayDemo.SetColor(color.Green)
					arrayDemo.Render([]int{}, []int{i, i - 1}, alreadySortedIndexes, swapDelay)
				}
				if i-1 == leftIndex {
					alreadySortedIndexes = append(alreadySortedIndexes, leftIndex)
					arrayDemo.Render([]int{i}, []int{}, alreadySortedIndexes, indexDelay)
				}
			}

			leftIndex += 1
		}
	}

	return sample
}
