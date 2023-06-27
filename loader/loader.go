package loader

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type ResourceLoader struct {
}

func (r *ResourceLoader) GetSprite(filename string) *ebiten.Image {
	path := "misc/"

	img, _, err := ebitenutil.NewImageFromFile(path + filename)
	if err != nil {
		panic(err)
	}
	return img
}
