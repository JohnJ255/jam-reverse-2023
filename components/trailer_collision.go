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
	tc := &TrailerCollision{
		Component: framework.InitComponent(),
		Collision: framework.InitCollision(box),
	}
	tc.Collision.BehaviourOnCollide = tc.OnCollide
	return tc
}

func (c *TrailerCollision) SetOwner(obj framework.Updating) {
	c.Component.SetOwner(obj)
	c.Collision.SetEntity(obj.(*entities.TrailerEntity))
}

func (c *TrailerCollision) Start(f *framework.Framework) {
	c.f = f
	c.Collision.Start(f)
}

func (c *TrailerCollision) OnCollide(collide *framework.Collide) {
	trailer := c.GetOwner().(*entities.TrailerEntity)
	traktor := collide.Collision.GetEntity().(*entities.CarEntity)
	for _, cs := range collide.Contacts {
		if trailer.Trailer.Traktor != nil && trailer.Trailer.Traktor == traktor.Car {
			c.onCollideWithTractor(trailer, cs)
		} else {
			c.Collision.OnCollide(collide)
			if trailer.Trailer.Traktor != nil {
				trailer.Trailer.Traktor.OnTrailerContacts(cs)
			}
		}
	}
}

func (c *TrailerCollision) onCollideWithTractor(trailer *entities.TrailerEntity, cs framework.ContactSet) {
	sign := framework.Radian(0.07)
	if cs.MoveOut.ToRadian().LefterThan(trailer.Trailer.Position.Angle) {
		sign = -0.07
	}
	trailer.Trailer.Position.Angle += sign

	tlp := trailer.Trailer.GetTowbarLocalPosition()
	towbarPos := trailer.Trailer.Traktor.GetTowbarPosition()
	trailer.Trailer.Position.X = towbarPos.X - tlp.X
	trailer.Trailer.Position.Y = towbarPos.Y - tlp.Y
}
