package components

import (
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
)

type TrailerCollision struct {
	*framework.Component
	*framework.Collision
	obj          framework.ICollisionOwner
	carCollision *CarCollision
	f            *framework.Framework
}

func (c *TrailerCollision) GetName() string {
	return "TrailerCollision"
}

func NewTrailerCollision(obj framework.ICollisionSizableOwner, carComponent framework.IComponent) *TrailerCollision {
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
		Component:    framework.InitComponent(),
		Collision:    framework.InitCollision(box),
		carCollision: carComponent.(*CarCollision),
	}
}

func (c *TrailerCollision) Start(f *framework.Framework) {
	c.f = f
}

func (c *TrailerCollision) Update(dt float64) {
	cs := c.Collision.Intersect(c.carCollision.Collision)
	c.f.Debug.SetDebugDraw("TrailerCollision", c.f.Debug.DefaultDrawIntersections(cs))
	if len(cs) > 0 && cs[0].MoveOut != nil {
		trailer := c.GetOwner().(*entities.TrailerEntity)
		sign := framework.Radian(0.1)
		if cs[0].MoveOut.ToRadian().LefterThan(trailer.Trailer.Position.Angle) {
			sign = -0.1
		}
		trailer.Trailer.Position.Angle += sign
	}
}
