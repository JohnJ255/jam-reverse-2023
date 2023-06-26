package components

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
)

type PlayerCarControl struct {
	*framework.Component
	levelSize framework.Size
}

func (c *PlayerCarControl) GetName() string {
	return "PlayerCarControl"
}

func NewPlayerCarControl(levelSize framework.Size) *PlayerCarControl {
	return &PlayerCarControl{
		levelSize: levelSize,
		Component: framework.InitComponent(),
	}
}

func (c *PlayerCarControl) Start(f *framework.Framework) {
}

func (c *PlayerCarControl) Update(dt float64) {
	entity := c.GetOwner().(*entities.CarEntity)
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
	if ebiten.IsKeyPressed(ebiten.KeyR) {
		entity.Car.TowbarToggle()
	}

	entity.Car.Control(accelerate, wheel)

	entity.Car.Position.X = framework.Limited(entity.Car.Position.X, 0, c.levelSize.Length)
	entity.Car.Position.Y = framework.Limited(entity.Car.Position.Y, 0, c.levelSize.Height)
}
