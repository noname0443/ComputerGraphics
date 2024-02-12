package main

import (
	"image/color"
	_ "image/jpeg"
	"strconv"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

type UIData struct {
	MovementSpeed []int
	RotateSpeed   []int
}

func CreateBar(clr color.Color, Move, Rotate *int) *widget.Container {
	innerContainer := CreateContainer(clr)

	innerContainer2 := CreateContainer(clr)
	textRotate1 := CreateText("0")
	innerContainer2.AddChild(CreateButton("Rotate Up", func(args *widget.ButtonClickedEventArgs) {
		*Rotate += 1
		textRotate1.Label = strconv.Itoa(*Rotate)
	}))
	innerContainer2.AddChild(CreateButton("Rotate Down", func(args *widget.ButtonClickedEventArgs) {
		*Rotate -= 1
		textRotate1.Label = strconv.Itoa(*Rotate)
	}))
	innerContainer2.AddChild(textRotate1)

	innerContainer.AddChild(innerContainer2)

	innerContainer1 := CreateContainer(clr)
	textMove1 := CreateText("0")
	innerContainer1.AddChild(CreateButton("Speed Up", func(args *widget.ButtonClickedEventArgs) {
		*Move += 1
		textMove1.Label = strconv.Itoa(*Move)
	}))
	innerContainer1.AddChild(CreateButton("Speed Down", func(args *widget.ButtonClickedEventArgs) {
		*Move -= 1
		textMove1.Label = strconv.Itoa(*Move)
	}))
	innerContainer1.AddChild(textMove1)
	innerContainer.AddChild(innerContainer1)

	return innerContainer
}

func initUI(data UIData) *widget.Container {
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(1),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(30)),
		)),
	)

	innerContainer := CreateContainer(color.RGBA{0, 0, 0, 255})

	innerContainer2 := CreateContainer(color.RGBA{0, 0, 0, 255})
	textRotate1 := CreateText("0")
	innerContainer2.AddChild(CreateButton("Rotate Up", func(args *widget.ButtonClickedEventArgs) {
		data.RotateSpeed[0] += 1
		textRotate1.Label = strconv.Itoa(data.RotateSpeed[0])
	}))
	innerContainer2.AddChild(CreateButton("Rotate Down", func(args *widget.ButtonClickedEventArgs) {
		data.RotateSpeed[0] -= 1
		textRotate1.Label = strconv.Itoa(data.RotateSpeed[0])
	}))
	innerContainer2.AddChild(textRotate1)
	innerContainer.AddChild(innerContainer2)

	rootContainer.AddChild(innerContainer)
	rootContainer.AddChild(CreateBar(color.RGBA{0, 0, 255, 255}, &data.MovementSpeed[0], &data.RotateSpeed[1]))
	rootContainer.AddChild(CreateBar(color.RGBA{0, 255, 0, 255}, &data.MovementSpeed[1], &data.RotateSpeed[2]))
	rootContainer.AddChild(CreateBar(color.RGBA{255, 0, 0, 255}, &data.MovementSpeed[2], &data.RotateSpeed[3]))
	rootContainer.AddChild(CreateBar(color.RGBA{255, 0, 255, 255}, &data.MovementSpeed[3], &data.RotateSpeed[4]))

	return rootContainer
}
