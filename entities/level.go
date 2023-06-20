package entities

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/loader"
	"strconv"
)

type Level struct {
	name            string
	backgroundImage *ebiten.Image
}

func NewLevel(index int) *Level {
	res := &loader.ResourceLoader{}
	return &Level{
		name:            "level " + strconv.Itoa(index),
		backgroundImage: res.GetSprite(fmt.Sprintf("level%d.png", index)),
	}
}

func (l *Level) GetSprite() *ebiten.Image {
	return l.backgroundImage
}

func (l *Level) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	return &ebiten.DrawImageOptions{}
}
