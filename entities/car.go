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
		Sprite:   framework.InitSprites(),
		Car:      car,
		IsPlayer: ct == framework.Player,
	}
	c.GameEntity = framework.InitGameEntity(c)
	c.LoadResources(&loader.ResourceLoader{}, loader.CarFileNames[ct])

	return c
}

func (c *CarEntity) GetPivot() framework.VecUV {
	return c.Car.Pivot
}

func (c *CarEntity) GetSize() framework.Size {
	return c.Car.Size
}

func (c *CarEntity) GetPosition() framework.Vec2 {
	return c.Car.Position.Vec2
}

func (c *CarEntity) GetRotation() framework.Radian {
	return c.Car.Position.Angle
}

func (c *CarEntity) GetScale() framework.Vec2 {
	return framework.Vec2{1, 1}
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

func (c *CarEntity) GetName() string {
	return "car"
}

func (c *CarEntity) GetModel() framework.Model {
	return c.Car
}
