package main

import (
	"image/color"
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ellipses []*Ellipse
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

var i int

func (g *Game) Draw(screen *ebiten.Image) {
	if len(g.ellipses) == 0 {
		ellipse := NewEllipse(400, 200, color.RGBA{0, 0, 255, 255})
		ellipse2 := NewEllipse(200, 100, color.RGBA{255, 0, 0, 255})
		ellipse3 := NewEllipse(100, 50, color.RGBA{0, 255, 0, 255})
		ellipse4 := NewEllipse(50, 25, color.RGBA{255, 255, 255, 255})
		ellipse5 := NewEllipse(25, 12, color.RGBA{255, 0, 255, 255})

		ellipse.Attach(ellipse2)
		ellipse2.Attach(ellipse3)
		ellipse3.Attach(ellipse4)
		ellipse4.Attach(ellipse5)

		g.ellipses = []*Ellipse{ellipse, ellipse2, ellipse3, ellipse4, ellipse5}
	}

	screenWidth, screenHeight := ebiten.WindowSize()
	g.ellipses[0].Move(screenWidth/2, screenHeight/2, screen)
	g.ellipses[0].IterateMove(3)
	g.ellipses[1].IterateMove(6)
	g.ellipses[2].IterateMove(9)
	g.ellipses[3].IterateMove(12)

	g.ellipses[0].IterateRotate(3)
	g.ellipses[1].IterateRotate(6)
	g.ellipses[2].IterateRotate(9)
	g.ellipses[3].IterateRotate(12)
	g.ellipses[4].IterateRotate(15)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	ebiten.SetWindowTitle("Ellipse Drawer")
	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
