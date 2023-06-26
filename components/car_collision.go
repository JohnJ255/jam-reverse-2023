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
	cc := &CarCollision{
		Component: framework.InitComponent(),
		Collision: framework.InitCollision(framework.NewBoxCollision(obj.GetSize(), obj)),
	}
	cc.Collision.BehaviourOnCollide = cc.OnCollide
	return cc
}

func (c *CarCollision) SetOwner(obj framework.Updating) {
	c.Component.SetOwner(obj)
	c.Collision.SetEntity(obj.(*entities.CarEntity))
}

func (c *CarCollision) GetName() string {
	return "CarCollision"
}

func (c *CarCollision) OnCollide(collide *framework.Collide) {
	car := c.GetOwner().(*entities.CarEntity)
	if car.Car.Trailer != nil && collide.Collision.GetEntity().GetPosition() == car.Car.Trailer.(*models.Trailer).Position.Vec2 {
		return
	}
	c.Collision.OnCollide(collide)

	for _, cs := range collide.Contacts {
		sign := framework.Radian(0.01)
		if car.GetRotation().LefterThan((*cs.MoveOut).ToRadian()) {
			sign = -0.01
		}
		car.SetRotation(car.GetRotation() + sign)
	}
}
