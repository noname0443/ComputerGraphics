package main

import (
	_ "image/jpeg"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screenWidth, screenHeight := ebiten.WindowSize()

	//wc := sync.WaitGroup{}

	ellipse := ebiten.NewImage(screenWidth, screenHeight)
	DrawEllipse(100, 50, screenWidth/2, screenHeight/2, 90, ellipse)

	screen.DrawImage(ellipse, nil)
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
