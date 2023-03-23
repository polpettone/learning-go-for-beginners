package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 200
	screenHeight = 200
)

const IMAGE_FILE_PATH = "deskpet2.gif"

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Move an image")
	ebiten.SetWindowPosition(1000, 1000)

	game := &Game{
		Pet:   NewPet(50, 50, IMAGE_FILE_PATH),
		X:     500,
		Y:     500,
		Clock: 0,
	}
	ebiten.SetScreenTransparent(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
