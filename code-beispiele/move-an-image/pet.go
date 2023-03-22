package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Pet struct {
	X     int
	Y     int
	Image *ebiten.Image
}

func NewPet(x, y int, imagePath string) *Pet {

	img, err := loadImage(imagePath)
	if err != nil {
		log.Fatal(err)
	}

	return &Pet{
		X:     x,
		Y:     y,
		Image: img,
	}
}

func (p *Pet) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(p.X), float64(p.Y))
	screen.DrawImage(p.Image, op)

}

func loadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	return img, nil
}
