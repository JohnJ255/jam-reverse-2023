package components

import (
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
)

type TrailerCollision struct {
	*framework.Component
	*framework.Collision
	f *framework.Framework
}

func (c *TrailerCollision) GetName() string {
	return "TrailerCollision"
}

func NewTrailerCollision(obj framework.ICollisionComponentOwner) *TrailerCollision {
	points := []framework.VecUV{
		{0, 0},
		{0.7, 0},
		{0.7, 0.15},
		{0.95, 0.5},
		{0.7, 0.85},
		{0.7, 1},
		{0, 1},
	}
	box := framework.NewPolygonCollisionUV(points, obj.GetSize(), obj)
	return &TrailerCollision{
		Component: framework.InitComponent(),
		Collision: framework.InitCollision(box),
	}
}

func (c *TrailerCollision) SetOwner(obj framework.Updating) {
	c.Component.SetOwner(obj)
	c.Collision.SetEntity(obj.(*entities.TrailerEntity))
}

func (c *TrailerCollision) Start(f *framework.Framework) {
	c.f = f
	f.RegisterCollision(c.Collision, c.GetOwner().(*entities.TrailerEntity))
}

func (c *TrailerCollision) Update(_ float64) {
	for _, collision := range c.f.GetClosestCollisonsFor(c.Collision) {
		cs := c.Collision.Intersect(collision)
		if len(cs) > 0 && cs[0].MoveOut != nil {
			trailer := c.GetOwner().(*entities.TrailerEntity)
			traktor := collision.GetEntity().(*entities.CarEntity)
			if trailer.Trailer.Traktor != nil && trailer.Trailer.Traktor == traktor.Car {
				c.onCollideWithTractor(trailer, cs)
			} else {
				c.OnCollide(trailer, cs)
			}
		}
	}
}

func (c *TrailerCollision) onCollideWithTractor(trailer *entities.TrailerEntity, cs []framework.ContactSet) {
	sign := framework.Radian(0.1)
	if cs[0].MoveOut.ToRadian().LefterThan(trailer.Trailer.Position.Angle) {
		sign = -0.1
	}
	trailer.Trailer.Position.Angle += sign
}

func (c *TrailerCollision) OnCollide(trailer *entities.TrailerEntity, cs []framework.ContactSet) {
	cts := cs[0]
	if cts.MoveOut == nil {
		return
	}
	if trailer.Trailer.Traktor == nil {
		sign := framework.Radian(0.01)
		if trailer.Trailer.Position.Angle.LefterThan((*cts.MoveOut).ToRadian()) {
			sign = -0.01
		}
		trailer.Trailer.Position.Angle += sign
	}
	trailer.Trailer.Position.X += cs[0].MoveOut.X
	trailer.Trailer.Position.Y += cs[0].MoveOut.Y
	if trailer.Trailer.Traktor != nil {
		trailer.Trailer.Traktor.OnTrailerContacts(cs)
	}
}
