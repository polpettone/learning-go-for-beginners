package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Pet struct {
	X                int
	Y                int
	Images           []*ebiten.Image
	CurrentImage     *ebiten.Image
	AnimationCounter int
}

func NewPet(x, y int, imagePath []string) *Pet {

	images := []*ebiten.Image{}
	for _, v := range imagePath {
		img, err := loadImage(v)
		if err != nil {
			log.Fatal(err)
		}
		images = append(images, img)
	}

	return &Pet{
		X:                x,
		Y:                y,
		Images:           images,
		AnimationCounter: 0,
		CurrentImage:     images[0],
	}
}

func (p *Pet) Draw(screen *ebiten.Image, gameClock int) {

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.3, 0.3)
	op.GeoM.Translate(float64(p.X), float64(p.Y))

	screen.DrawImage(p.CurrentImage, op)

	if gameClock%20 == 0 {
		p.CurrentImage = p.Images[p.AnimationCounter]
		p.AnimationCounter = p.AnimationCounter + 1
		if p.AnimationCounter >= 3 {
			p.AnimationCounter = 0
		}
	}

}

func loadImage(path string) (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	return img, nil
}
