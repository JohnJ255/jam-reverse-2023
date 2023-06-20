package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/loader"
	"reverse-jam-2023/models"
)

type CarEntity struct {
	*framework.GameEntity
	*framework.Sprite
	Car      *models.Car
	IsPlayer bool
}

func NewCar(ct framework.ControlType, car *models.Car) *CarEntity {
	c := &CarEntity{
		GameEntity: framework.InitGameEntity(),
		Sprite:     framework.InitSprites(),
		Car:        car,
		IsPlayer:   ct == framework.Player,
	}
	c.LoadResources(&loader.ResourceLoader{}, loader.CarFileNames[ct])

	return c
}

func (c *CarEntity) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	op := c.PivotTransform(scale, c.Car.Size, c.Car.Pivot)
	op.GeoM.Rotate(float64(c.Car.Position.Angle))
	op.GeoM.Translate(c.Car.Position.X, c.Car.Position.Y)

	return op
}

func (c *CarEntity) AddComponent(comp framework.IComponent) {
	comp.SetOwner(c)
	c.GameEntity.AddComponent(comp)
}
