package components

import (
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
)

type CarCollision struct {
	*framework.Component
	*framework.Collision
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
	f.RegisterCollision(c.Collision, c.GetOwner().(*entities.CarEntity))
}

func (c *CarCollision) Update(dt float64) {
}
