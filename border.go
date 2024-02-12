package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type phase string

const (
	blackPhase phase = "black"
	whitePhase phase = "white"
)

type Point struct {
	x int
	y int
}

func GetBorder(img *ebiten.Image) [][]int {
	border := [][]int{}
	borderMap := map[Point]bool{}

	startX, startY := -1, -1
	for i := 0; i < img.Bounds().Dx(); i++ {
		flag := false
		for j := 0; j < img.Bounds().Dy(); j++ {
			if (img.At(i, j) != color.RGBA{}) {
				startX, startY = i, j
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	border = append(border, []int{startX, startY})
	black := Point{startX, startY}

	isValidBlackPixel := func(x, y int) bool {
		ok := borderMap[Point{x, y}]
		if (img.At(x, y) != color.RGBA{}) && !ok {
			return true
		}
		return false
	}

	isWhitePixel := func(x, y int) bool {
		return img.At(x, y) == color.RGBA{}
	}

	Points := []Point{
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
	}

	flag := true

	for flag {
		flag = false

		for i := 0; i < len(Points); i++ {
			x, y := Points[(i+1)%len(Points)].x, Points[(i+1)%len(Points)].y
			xPrev, yPrev := Points[i].x, Points[i].y

			if isWhitePixel(black.x+xPrev, black.y+yPrev) && isValidBlackPixel(black.x+x, black.y+y) {
				flag = true
				border = append(border, []int{black.x, black.y})
				borderMap[black] = true
				black = Point{black.x + x, black.y + y}
			}
		}
	}

	return border
}

func DrawPixels(pixels [][]int, screen *ebiten.Image) {
	for _, v := range pixels {
		screen.Set(v[0], v[1], color.RGBA{255, 0, 0, 255})
	}
}
