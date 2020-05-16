package sorting

import (
	"github.com/gookit/color"
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/demo"

	"github.com/gosuri/uilive"
)

func BubbleSortDemo(inputArray []int, cmpClause int8, indexDelay time.Duration, swapDelay time.Duration) []int {
	sample := make([]int, len(inputArray))
	copy(sample, inputArray)

	arrayDemo := demo.NewArrayDemo(sample, uilive.New())

	if cmpClause == 0 {
		return sample
	}

	alreadySortedIndexes := make([]int, 0)

	if cmpClause < 0 {
		for i := 0; i < (len(sample) - 1); i++ {
			for j := 0; j < (len(sample) - i - 1); j++ {
				arrayDemo.Render([]int{j}, []int{}, alreadySortedIndexes, indexDelay)

				if sample[j] > sample[j+1] {
					arrayDemo.SetColor(color.Red)
					arrayDemo.Render([]int{}, []int{j, j + 1}, alreadySortedIndexes, swapDelay)

					sample[j], sample[j+1] = sample[j+1], sample[j]

					arrayDemo.SetColor(color.Green)
					arrayDemo.Render([]int{}, []int{j, j + 1}, alreadySortedIndexes, swapDelay)
				}
				if j+1 == (len(sample) - i - 1) {
					alreadySortedIndexes = append(alreadySortedIndexes, j+1)
				}
				if i+1 == (len(sample) - 1) {
					arrayDemo.Render([]int{j}, []int{}, alreadySortedIndexes, indexDelay)
				}
			}
		}
	}
	if cmpClause > 0 {
		for i := 0; i < (len(sample) - 1); i++ {
			for j := 0; j < (len(sample) - i - 1); j++ {
				arrayDemo.Render([]int{j}, []int{}, alreadySortedIndexes, indexDelay)

				if sample[j] < sample[j+1] {
					arrayDemo.SetColor(color.Red)
					arrayDemo.Render([]int{}, []int{j, j + 1}, alreadySortedIndexes, swapDelay)

					sample[j], sample[j+1] = sample[j+1], sample[j]

					arrayDemo.SetColor(color.Green)
					arrayDemo.Render([]int{}, []int{j, j + 1}, alreadySortedIndexes, swapDelay)
				}
				if j+1 == (len(sample) - i - 1) {
					alreadySortedIndexes = append(alreadySortedIndexes, j+1)
				}
				if i+1 == (len(sample) - 1) {
					arrayDemo.Render([]int{j}, []int{}, alreadySortedIndexes, indexDelay)
				}
			}
		}
	}

	arrayDemo.Close()

	return sample
}
