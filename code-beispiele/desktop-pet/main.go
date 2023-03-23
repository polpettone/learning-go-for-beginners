package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 200
	screenHeight = 200
)

func main() {

	catImages := []string{
		"assets/pet-1/pet1-1.png",
		"assets/pet-1/pet1-2.png",
		"assets/pet-1/pet1-3.png"}

	hamsterImages := []string{
		"assets/pet-2/hamster-1.png",
		"assets/pet-2/hamster-2.png",
		"assets/pet-2/hamster-3.png",
		"assets/pet-2/hamster-4.png"}

	cat := NewPet(50, 50, 0.3, catImages)
	hamster := NewPet(50, 50, 0.8, hamsterImages)

	pets := []*Pet{cat, hamster}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Desktop Pet")
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetWindowPosition(1000, 1000)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)

	game := &Game{
		Pet:   pets[0],
		X:     500,
		Y:     500,
		Clock: 0,
	}
	ebiten.SetScreenTransparent(true)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
