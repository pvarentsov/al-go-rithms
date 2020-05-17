package sorting

import (
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/demo"

	"github.com/gookit/color"
	"github.com/gosuri/uilive"
)

func CombSortDemo(inputArray []int, cmpClause int8, indexDelay time.Duration, swapDelay time.Duration) []int {
	sample := make([]int, len(inputArray))
	copy(sample, inputArray)

	if cmpClause == 0 {
		return sample
	}

	const factor = 1.247
	gap := len(inputArray) - 1

	arrayDemo := demo.NewArrayDemo(sample, uilive.New())

	for gap >= 1 {
		if cmpClause < 0 {
			for i := 0; i+gap < len(inputArray); i++ {
				arrayDemo.Render([]int{i, i + gap}, []int{}, []int{}, indexDelay)

				if sample[i] > sample[i+gap] {
					arrayDemo.SetColor(color.Red)
					arrayDemo.Render([]int{}, []int{i, i + gap}, []int{}, swapDelay)

					sample[i], sample[i+gap] = sample[i+gap], sample[i]

					arrayDemo.SetColor(color.Green)
					arrayDemo.Render([]int{}, []int{i, i + gap}, []int{}, swapDelay)
				}
			}
		}
		if cmpClause > 0 {
			for i := 0; i+gap < len(inputArray); i++ {
				arrayDemo.Render([]int{i, i + gap}, []int{}, []int{}, indexDelay)

				if sample[i] < sample[i+gap] {
					arrayDemo.SetColor(color.Red)
					arrayDemo.Render([]int{}, []int{i, i + gap}, []int{}, swapDelay)

					sample[i], sample[i+gap] = sample[i+gap], sample[i]

					arrayDemo.SetColor(color.Green)
					arrayDemo.Render([]int{}, []int{i, i + gap}, []int{}, swapDelay)
				}
			}
		}

		gap = int(float32(gap) / factor)

		if gap <= 1 {
			arrayDemo.Render([]int{len(inputArray) - 1}, []int{}, []int{}, indexDelay)
		}
	}

	return sample
}
