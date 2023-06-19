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
	Car *models.Car
}

func NewCar(ct framework.ControlType, car *models.Car) *CarEntity {
	c := &CarEntity{
		SpriteEntity: framework.InitSprites(-math.Pi / 2),
		Car:          car,
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

func (c *CarEntity) Control(params interface{}) {
	var accelerate float64
	var wheel float64
	switch v := params.(type) {
	case map[string]float64:
		accelerate = v["accelerate"]
		wheel = v["wheel"]
	}

	c.Car.Control(accelerate, wheel)
}
