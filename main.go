package main

import (
	"image/color"
	_ "image/jpeg"
	"log"

	"github.com/ebitenui/ebitenui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	ellipses []*Ellipse
	ui       *ebitenui.UI

	MovementSpeed []int
	RotateSpeed   []int
}

func NewGame() *Game {
	move := make([]int, 5)
	rotate := make([]int, 5)

	return &Game{
		ui: &ebitenui.UI{
			Container: initUI(UIData{move, rotate}),
		},
		MovementSpeed: move,
		RotateSpeed:   rotate,
	}
}

func (g *Game) Update() error {
	g.ui.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.ui.Draw(screen)

	if len(g.ellipses) == 0 {
		ellipse := NewEllipse(400, 200, color.RGBA{0, 0, 0, 255})
		ellipse2 := NewEllipse(200, 100, color.RGBA{0, 0, 255, 255})
		ellipse3 := NewEllipse(100, 50, color.RGBA{0, 255, 0, 255})
		ellipse4 := NewEllipse(50, 25, color.RGBA{255, 0, 0, 255})
		ellipse5 := NewEllipse(25, 12, color.RGBA{255, 0, 255, 255})

		ellipse.Attach(ellipse2)
		ellipse2.Attach(ellipse3)
		ellipse3.Attach(ellipse4)
		ellipse4.Attach(ellipse5)

		g.ellipses = []*Ellipse{ellipse, ellipse2, ellipse3, ellipse4, ellipse5}
	}

	screenWidth, screenHeight := ebiten.WindowSize()
	g.ellipses[0].Move(screenWidth*3/4, screenHeight/2, screen)

	for i := 0; i < 5; i++ {
		if i != 4 {
			g.ellipses[i].IterateMove(g.MovementSpeed[i])
		}
		g.ellipses[i].IterateRotate(g.RotateSpeed[i])
	}
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
