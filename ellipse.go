package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func Ellipse(width, height int) *ebiten.Image {
	m := width
	if m < height {
		m = height
	}
	img := ebiten.NewImage(width, height)

	width /= 2
	height /= 2

	for y := -height; y <= height; y++ {
		for x := -width; x <= width; x++ {
			if x*x*height*height+y*y*width*width <= height*height*width*width {
				img.Set(width+x, height+y, color.White)
			}
		}
	}
	return img
}

func Move(x, y int, dst, tgt *ebiten.Image) {
	for i := 0; i < tgt.Bounds().Dx(); i++ {
		for j := 0; j < tgt.Bounds().Dy(); j++ {
			dst.Set(i+x, y+j, tgt.At(i, j))
		}
	}
}

func Rotate(degree int, img *ebiten.Image) *ebiten.Image {
	m := img.Bounds().Dy()
	if m < img.Bounds().Dx() {
		m = img.Bounds().Dx()
	}

	newImg := ebiten.NewImage(m, m)
	rad := math.Pi * float64(degree) / 180.0

	for i := 0; i < img.Bounds().Dx(); i++ {
		for j := 0; j < img.Bounds().Dy(); j++ {
			x := float64(i - img.Bounds().Dx()/2)
			y := float64(j - img.Bounds().Dy()/2)

			I := int(x*math.Cos(rad)+y*-math.Sin(rad)) + newImg.Bounds().Dx()/2
			J := int(x*math.Sin(rad)+y*math.Cos(rad)) + newImg.Bounds().Dy()/2

			newImg.Set(I, J, img.At(i, j))
			newImg.Set(I-1, J, img.At(i, j))
			newImg.Set(I, J-1, img.At(i, j))
			newImg.Set(I-1, J-1, img.At(i, j))
		}
	}

	return newImg
}

func GetCenter(img *ebiten.Image) (int, int) {
	X := 0
	for i := 0; i < img.Bounds().Dx(); i++ {
		for j := 0; j < img.Bounds().Dy(); j++ {
			if img.At(i, j) != color.Transparent {
				X += i
			}
		}
	}

	Y := 0
	for i := 0; i < img.Bounds().Dy(); i++ {
		for j := 0; j < img.Bounds().Dx(); j++ {
			if img.At(i, j) != color.Transparent {
				Y += i
			}
		}
	}

	Square := img.Bounds().Dx() * img.Bounds().Dy()
	return X / Square, Y / Square
}

func DrawEllipse(width, height int, x, y int, degree int, screen *ebiten.Image) {
	ellipse := Ellipse(width, height)
	ellipse = Rotate(degree, ellipse)
	w, h := GetCenter(ellipse)
	Move(x-w, y-h, screen, ellipse)
}
