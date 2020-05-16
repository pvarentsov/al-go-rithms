package sorting

import (
	"github.com/gookit/color"
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/demo"

	"github.com/gosuri/uilive"
)

func StupidSortDemo(inputArray []int, cmpClause int8, indexDelay time.Duration, swapDelay time.Duration) []int {
	sample := make([]int, len(inputArray))
	copy(sample, inputArray)

	if cmpClause == 0 {
		return sample
	}

	arrayDemo := demo.NewArrayDemo(sample, uilive.New())

	index := 0

	arrayDemo.Render([]int{0}, []int{}, []int{}, indexDelay)

	for index < (len(sample) - 1) {
		if cmpClause < 0 {
			if sample[index] < sample[index+1] {
				index++
				arrayDemo.Render([]int{index}, []int{}, []int{}, indexDelay)
			} else {
				arrayDemo.SetColor(color.Red)
				arrayDemo.Render([]int{}, []int{index, index + 1}, []int{}, swapDelay)

				sample[index], sample[index+1] = sample[index+1], sample[index]

				arrayDemo.SetColor(color.Green)
				arrayDemo.Render([]int{}, []int{index, index + 1}, []int{}, swapDelay)

				index = 0
				arrayDemo.Render([]int{index}, []int{}, []int{}, indexDelay)
			}
		}
		if cmpClause > 0 {
			if sample[index] > sample[index+1] {
				index++
				arrayDemo.Render([]int{index}, []int{}, []int{}, indexDelay)
			} else {
				arrayDemo.SetColor(color.Red)
				arrayDemo.Render([]int{}, []int{index, index + 1}, []int{}, swapDelay)

				sample[index], sample[index+1] = sample[index+1], sample[index]

				arrayDemo.SetColor(color.Green)
				arrayDemo.Render([]int{}, []int{index, index + 1}, []int{}, swapDelay)

				index = 0
				arrayDemo.Render([]int{index}, []int{}, []int{}, indexDelay)
			}
		}
	}

	arrayDemo.Close()

	return sample
}
