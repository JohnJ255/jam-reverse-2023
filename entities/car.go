package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/loader"
	"reverse-jam-2023/models"
)

type CarEntity struct {
	Car       *models.Car
	imgs      []*ebiten.Image
	DrawAngle float64
}

func NewCar(ct framework.ControlType, car *models.Car) *CarEntity {
	c := &CarEntity{
		imgs:      make([]*ebiten.Image, 0, len(loader.FileNames[ct])),
		Car:       car,
		DrawAngle: -math.Pi / 2,
	}
	res := &loader.Resource{}
	for _, fileName := range loader.FileNames[ct] {
		img := res.GetSprite(fileName)
		c.imgs = append(c.imgs, img)
	}

	return c
}

func (c *CarEntity) GetSprite() *ebiten.Image {
	return c.imgs[0]
}

func (c *CarEntity) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(-c.Car.Size.Width/2, -c.Car.Size.Length/2)
	op.GeoM.Rotate(c.Car.Position.Angle - c.DrawAngle)
	op.GeoM.Translate(c.Car.Position.X+100, c.Car.Position.Y+100)

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
