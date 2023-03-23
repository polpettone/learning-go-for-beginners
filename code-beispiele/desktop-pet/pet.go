package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Pet struct {
	X                int
	Y                int
	Images           []*ebiten.Image
	CurrentImage     *ebiten.Image
	AnimationCounter int
	Scale            float64
}

func NewPet(x, y int, scale float64, path string) *Pet {

	images, err := loadImages(path)

	if err != nil {
		log.Fatal(err)
	}

	return &Pet{
		X:                x,
		Y:                y,
		Images:           images,
		AnimationCounter: 0,
		CurrentImage:     images[0],
		Scale:            scale,
	}
}

func (p *Pet) Draw(screen *ebiten.Image, gameClock int) {

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(p.Scale, p.Scale)
	op.GeoM.Translate(float64(p.X), float64(p.Y))

	screen.DrawImage(p.CurrentImage, op)

	if gameClock%40 == 0 {
		p.CurrentImage = p.Images[p.AnimationCounter]
		p.AnimationCounter = p.AnimationCounter + 1
		if p.AnimationCounter >= len(p.Images) {
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

func loadImages(path string) ([]*ebiten.Image, error) {
	images := []*ebiten.Image{}

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		img, _, err := ebitenutil.NewImageFromFile(path+"/"+file.Name(), ebiten.FilterDefault)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}

	return images, nil
}
