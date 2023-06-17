package loader

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"reverse-jam-2023/framework"
)

var FileNames = map[framework.ControlType][]string{
	framework.Player: {
		"Car_1_01.png",
		"Car_1_02.png",
		"Car_1_03.png",
		"Car_1_04.png",
		"Car_1_05.png",
	},
	framework.Computer: {
		"Car_3_01.png",
		"Car_3_02.png",
		"Car_3_03.png",
		"Car_3_04.png",
		"Car_3_05.png",
	},
}

type Resource struct {
}

func (r *Resource) GetSprite(filename string) *ebiten.Image {
	path := "images/"

	img, _, err := ebitenutil.NewImageFromFile(path + filename)
	if err != nil {
		panic(err)
	}
	return img
}
