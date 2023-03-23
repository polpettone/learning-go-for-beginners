package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1280
	screenHeight = 640
)

const IMAGE_FILE_PATH = "gopher.png"

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Move an image")
	game := &Game{
		Pet: NewPet(100, 100, IMAGE_FILE_PATH),
	}
	ebiten.SetScreenTransparent(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
