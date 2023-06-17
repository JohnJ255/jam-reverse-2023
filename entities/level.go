package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/loader"
	"strconv"
)

type Level struct {
	name            string
	backgroundImage *ebiten.Image
}

func NewLevel(index int) *Level {
	res := &loader.Resource{}
	return &Level{
		name:            "level " + strconv.Itoa(index),
		backgroundImage: res.GetSprite("background.png"),
	}
}

func (l *Level) GetSprite() *ebiten.Image {
	return l.backgroundImage
}

func (l *Level) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	return &ebiten.DrawImageOptions{}
}
