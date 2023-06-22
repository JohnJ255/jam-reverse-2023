package components

import (
	"reverse-jam-2023/framework"
)

type CarCollision struct {
	*framework.Component
	*framework.Collision
	obj framework.ICollisionOwner
}

func NewCarCollision(obj framework.ICollisionSizableOwner) *CarCollision {
	return &CarCollision{
		Component: framework.InitComponent(),
		Collision: framework.InitCollision(framework.NewBoxCollision(obj.GetSize(), obj)),
	}
}

func (c *CarCollision) GetName() string {
	return "CarCollision"
}

func (c *CarCollision) Start(f *framework.Framework) {
}

func (c *CarCollision) Update(dt float64) {
}
