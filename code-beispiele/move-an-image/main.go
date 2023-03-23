package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 500
	screenHeight = 500
)

func main() {

	images := []string{"pet1-1.png", "pet1-2.png", "pet1-3.png"}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Move an image")
	ebiten.SetWindowPosition(1000, 1000)

	game := &Game{
		Pet:   NewPet(50, 50, images),
		X:     500,
		Y:     500,
		Clock: 0,
	}
	ebiten.SetScreenTransparent(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
