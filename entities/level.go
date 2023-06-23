package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/loader"
	"strconv"
)

type Level struct {
	*framework.Sprite
	name string
	size framework.Size
}

func NewLevel(index int) *Level {
	level := &Level{
		Sprite: framework.InitSprites(),
		name:   "level " + strconv.Itoa(index),
		size:   framework.Size{1200, 600},
	}
	level.LoadResources(&loader.ResourceLoader{}, loader.LevelFileNames[index])
	return level
}

func (l *Level) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	return l.PivotTransform(scale, l.size, framework.VecUV{})
}
