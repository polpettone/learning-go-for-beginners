package main

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Game struct {
	Pet   *Pet
	X     int
	Y     int
	Clock int
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func (g *Game) Update(screen *ebiten.Image) error {

	g.Clock = g.Clock + 1

	if g.Clock%60 == 0 {
		x, y := ebiten.WindowPosition()
		g.X = x
		g.Y = y
		g.X = g.X + random(-100, 100)
		g.Y = g.Y + random(-100, 100)
		ebiten.SetWindowPosition(g.X, g.Y)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.Pet.Y = g.Pet.Y - 10
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.Pet.Y = g.Pet.Y + 10
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.Pet.X = g.Pet.X + 10
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyA) {
		g.Pet.X = g.Pet.X - 10
	}

	return nil

}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Pet.Draw(screen, g.Clock)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
