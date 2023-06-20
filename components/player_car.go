package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
)

type PlayerCarControl struct {
	*framework.Component
}

func NewPlayerCarControl() *PlayerCarControl {
	return &PlayerCarControl{
		Component: framework.InitComponent(),
	}
}

func (c *PlayerCarControl) Start() {
}

func (c *PlayerCarControl) Update(dt float64) {
	accelerate := 0.0
	wheel := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		wheel = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		wheel = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		accelerate = 1.0
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		accelerate = -0.3
	}

	c.GetOwner().(*entities.CarEntity).Car.Control(accelerate, wheel)
}
