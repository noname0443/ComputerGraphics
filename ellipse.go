package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ellipse struct {
	image          *ebiten.Image
	width, height  int
	border         []Point
	attchedEllipse *Ellipse
	clr            color.Color
	iteratorMove   int
	iteratorRotate int
}

func NewEllipse(width, height int, clr color.Color) *Ellipse {
	m := width
	if height > m {
		m = height
	}
	img := ebiten.NewImage(m, m)

	width /= 2
	height /= 2

	for y := -height; y <= height; y++ {
		for x := -width; x <= width; x++ {
			if x*x*height*height+y*y*width*width <= height*height*width*width {
				img.Set(width+x, height+y, clr)
			}
		}
	}
	return &Ellipse{
		image:  img,
		width:  width,
		height: height,
		clr:    clr,
	}
}

func RotatePixel(point Point, degree int) Point {
	rad := math.Pi * float64(degree) / 180.0

	X := float64(point.x)
	Y := float64(point.y)

	return Point{int(X*math.Cos(rad) + Y*-math.Sin(rad)), int(X*math.Sin(rad) + Y*math.Cos(rad))}
}

func (el *Ellipse) Move(x, y int, destination *ebiten.Image) {
	x, y = x-el.image.Bounds().Dx()/2, y-el.image.Bounds().Dy()/2

	img := el.Rotate()
	for i := 0; i < img.Bounds().Dx(); i++ {
		for j := 0; j < img.Bounds().Dy(); j++ {
			if (img.At(i, j) != color.RGBA{}) {
				destination.Set(i+x, y+j, img.At(i, j))
			}
		}
	}

	if el.attchedEllipse != nil {
		point := el.border[el.iteratorMove]

		point = Point{point.x - el.width, point.y - el.height}
		point = RotatePixel(point, el.iteratorRotate)
		point = Point{point.x + img.Bounds().Dx()/2, point.y + img.Bounds().Dy()/2}

		el.attchedEllipse.Move(x+point.x, y+point.y, destination)
	}
}

func (el *Ellipse) IterateMove(speed int) {
	el.iteratorMove = (el.iteratorMove + speed) % len(el.border)
}

func (el *Ellipse) IterateRotate(speed int) {
	el.iteratorRotate = el.iteratorRotate + speed
}

func (el *Ellipse) Attach(ellipse *Ellipse) {
	el.EvaluateBorder()
	el.attchedEllipse = ellipse
}

func (el *Ellipse) Rotate() *ebiten.Image {
	m := el.image.Bounds().Dy()
	if m < el.image.Bounds().Dx() {
		m = el.image.Bounds().Dx()
	}

	newImg := ebiten.NewImage(m, m)
	rad := math.Pi * float64(el.iteratorRotate) / 180.0

	for i := 0; i < el.image.Bounds().Dx(); i++ {
		for j := 0; j < el.image.Bounds().Dy(); j++ {
			if (el.image.At(i, j) != color.RGBA{}) {
				x := float64(i - el.width)
				y := float64(j - el.height)

				I := int(x*math.Cos(rad)+y*-math.Sin(rad)) + newImg.Bounds().Dx()/2
				J := int(x*math.Sin(rad)+y*math.Cos(rad)) + newImg.Bounds().Dy()/2

				newImg.Set(I, J, el.image.At(i, j))
				newImg.Set(I-1, J, el.image.At(i, j))
				newImg.Set(I, J-1, el.image.At(i, j))
				newImg.Set(I-1, J-1, el.image.At(i, j))
			}
		}
	}

	return newImg
}

func (el *Ellipse) GetCenter() (int, int) {
	X := 0
	for i := 0; i < el.image.Bounds().Dx(); i++ {
		for j := 0; j < el.image.Bounds().Dy(); j++ {
			if el.image.At(i, j) != color.Transparent {
				X += i
			}
		}
	}

	Y := 0
	for i := 0; i < el.image.Bounds().Dy(); i++ {
		for j := 0; j < el.image.Bounds().Dx(); j++ {
			if el.image.At(i, j) != color.Transparent {
				Y += i
			}
		}
	}

	Square := el.image.Bounds().Dx() * el.image.Bounds().Dy()
	return X / Square, Y / Square
}

func (el *Ellipse) EvaluateBorder() {
	el.border = GetBorder(el.image)
}

func DrawEllipse(width, height int, x, y int, degree int, screen *ebiten.Image) {
	ellipse := NewEllipse(width, height, color.White)
	ellipse.Rotate()
	w, h := ellipse.GetCenter()
	ellipse.Move(x-w, y-h, screen)
}
