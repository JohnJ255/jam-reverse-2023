package framework

import "github.com/hajimehoshi/ebiten/v2"

type Drawing interface {
	GetSprite() *ebiten.Image
	GetTransforms(scale float64) *ebiten.DrawImageOptions
}

type Controlling interface {
	Drawing
	Control(params interface{})
}

type BaseEntity struct {
	Imgs      []*ebiten.Image
	DrawAngle float64
}
