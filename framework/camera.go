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
	return s.Background.PivotTransform(scale, VecUV{})
}

type FollowCamera struct {
	*StaticCamera
	pos Vec2
}

func NewFollowCamera(limiter ISizable, background *Sprite) *FollowCamera {
	return &FollowCamera{
		StaticCamera: &StaticCamera{
			Limiter:    limiter,
			Background: background,
		},
	}
}

func (f *FollowCamera) Control(obj IGameObject) {
	f.pos = obj.GetPosition()
}

func (f *FollowCamera) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	posUV := VecUV{f.pos.X, f.pos.Y}
	return f.Background.PivotTransform(scale, posUV)
}
