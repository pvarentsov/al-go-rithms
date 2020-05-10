package generator

import (
	"math/rand"
	"time"
)

func GenerateArray(size uint, shuffle bool) []int {
	array := make([]int, size)
	array[0] = 1

	for i := 1; i < len(array); i++ {
		array[i] = array[i-1] + 1
	}

	if shuffle {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(array), func(i, j int) { array[i], array[j] = array[j], array[i] })
	}

	return array
}
