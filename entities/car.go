package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
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
	op := &ebiten.DrawImageOptions{}
	spriteSize := image.Point{c.GetSprite().Bounds().Size().Y, c.GetSprite().Bounds().Size().X}
	scaleX := scale * c.Car.Size.Length / float64(spriteSize.X)
	scaleY := scale * c.Car.Size.Width / float64(spriteSize.Y)
	op.GeoM.Rotate(-c.DrawAngle)
	tx := c.Car.Size.Length * (1 - c.Car.Pivot.U)
	ty := -c.Car.Size.Width * c.Car.Pivot.V
	op.GeoM.Scale(scaleX, scaleY)
	op.GeoM.Translate(tx, ty)
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
