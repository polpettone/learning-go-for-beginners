package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Game struct {
	Pet *Pet
}

func (g *Game) Update(screen *ebiten.Image) error {

	log.Printf("%v", g.Pet)

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
	g.Pet.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
