package demo

import (
	"strings"
	"time"

	"github.com/pvarentsov/al-go-rithms/internal/converter"

	"github.com/gookit/color"
	"github.com/gosuri/uilive"
)

func NewArrayDemo(array []int, writer *uilive.Writer) *ArrayDemo {
	demo := ArrayDemo{
		array:  array,
		writer: writer,
		color:  color.Green,
	}

	demo.writer.Start()

	return &demo
}

type ArrayDemo struct {
	array  []int
	color  color.Color
	writer *uilive.Writer
}

func (demo *ArrayDemo) SetArray(array []int) {
	demo.array = array
}

func (demo *ArrayDemo) SetColor(color color.Color) {
	demo.color = color
}

func (demo *ArrayDemo) Render(underscoredIndexes []int, coloredIndexes []int, boldIndexes []int, delay time.Duration) {
	strArray := converter.IntArrayToStrArray(demo.array)

	for _, i := range underscoredIndexes {
		strArray[i] = color.OpUnderscore.Sprintf("%s", strArray[i])
	}
	for _, i := range coloredIndexes {
		strArray[i] = demo.color.Sprintf("%s", strArray[i])
	}
	for _, i := range boldIndexes {
		strArray[i] = color.Bold.Sprintf("%s", strArray[i])
	}

	color.Fprintf(demo.writer, "[ %s ]\n\n", strings.Join(strArray, " "))

	time.Sleep(time.Millisecond * delay)
}

func (demo *ArrayDemo) Close() {
	demo.writer.Stop()
}
