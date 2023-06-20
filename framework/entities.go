package framework

import "github.com/hajimehoshi/ebiten/v2"

type Drawing interface {
	GetSprite() *ebiten.Image
	GetTransforms(scale float64) *ebiten.DrawImageOptions
}

type Updating interface {
	Update(dt float64)
}

type Entity interface {
	Drawing
	Updating
}
