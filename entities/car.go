package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/loader"
	"reverse-jam-2023/models"
)

type CarEntity struct {
	*framework.SpriteEntity
	Car      *models.Car
	IsPlayer bool
}

func NewCar(ct framework.ControlType, car *models.Car) *CarEntity {
	c := &CarEntity{
		SpriteEntity: framework.InitSprites(math.Pi / 2),
		Car:          car,
		IsPlayer:     ct == framework.Player,
	}
	car.Position.X = 100
	car.Position.Y = 100
	c.LoadResources(&loader.Resource{}, loader.CarFileNames[ct])

	return c
}

func (c *CarEntity) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	op := c.PivotTransform(scale, c.Car.Size, c.Car.Pivot)
	op.GeoM.Rotate(c.Car.Position.Angle)
	op.GeoM.Translate(c.Car.Position.X, c.Car.Position.Y)

	return op
}

func (c *CarEntity) Control() {
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

	c.Car.Control(accelerate, wheel)
}

func (c *CarEntity) Update(dt float64) {
	if c.IsPlayer {
		c.Control()
	}
}
