package components

import (
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/models"
)

type CarCollision struct {
	*framework.Component
	*framework.Collision
	f *framework.Framework
}

func NewCarCollision(obj framework.ICollisionComponentOwner) *CarCollision {
	return &CarCollision{
		Component: framework.InitComponent(),
		Collision: framework.InitCollision(framework.NewBoxCollision(obj.GetSize(), obj)),
	}
}

func (c *CarCollision) GetName() string {
	return "CarCollision"
}

func (c *CarCollision) Start(f *framework.Framework) {
	c.f = f
	f.RegisterCollision(c.Collision, c.GetOwner().(*entities.CarEntity))
}

func (c *CarCollision) Update(dt float64) {
	for _, collision := range c.f.GetClosestCollisonsFor(c.Collision) {
		cs := c.Collision.Intersect(collision)
		if len(cs) > 0 && cs[0].MoveOut != nil {
			car := c.GetOwner().(*entities.CarEntity)
			if car.Car.Trailer != nil && collision.GetEntity().GetPosition() == car.Car.Trailer.(*models.Trailer).Position.Vec2 {
				continue
			}
			c.OnCollide(car, cs)
		}
	}
}

func (c *CarCollision) OnCollide(car *entities.CarEntity, cs []framework.ContactSet) {
	car.Car.Position.X += cs[0].MoveOut.X
	car.Car.Position.Y += cs[0].MoveOut.Y
}
