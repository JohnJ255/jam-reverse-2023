package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ICamera interface {
	Control(obj IGameObject)
	GetTransforms(scale float64) *ebiten.DrawImageOptions
}

//todo: make camera as GameEntity

type StaticCamera struct {
	Limiter    ISizable
	Background *Sprite
}

func NewStaticCamera(limiter ISizable, background *Sprite) *StaticCamera {
	return &StaticCamera{
		Limiter:    limiter,
		Background: background,
	}
}

func (s *StaticCamera) Control(_ IGameObject) {
}

func (s *StaticCamera) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	return s.Background.PivotTransform(scale, s.Limiter.GetSize(), VecUV{})
}

type FollowCamera struct {
	*StaticCamera
}
