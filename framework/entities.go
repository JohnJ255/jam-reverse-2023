package framework

import "github.com/hajimehoshi/ebiten/v2"

type Drawing interface {
	GetSprite() *ebiten.Image
	GetTransforms(scale float64) *ebiten.DrawImageOptions
}

type Entity interface {
	Drawing
	Update(dt float64)
}
