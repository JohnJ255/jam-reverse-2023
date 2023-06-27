package components

import (
	"reverse-jam-2023/framework"
)

type FixedCollision struct {
	*framework.Component
	*framework.Collision
	f *framework.Framework
}

func NewFixedCollision(obj framework.ICollisionComponentOwner) *FixedCollision {
	fc := &FixedCollision{
		Component: framework.InitComponent(),
		Collision: framework.InitCollision(framework.NewBoxCollision(obj.GetSize(), obj)),
	}
	return fc
}

func (fc *FixedCollision) GetName() string {
	return "FixedCollision"
}

func (fc *FixedCollision) SetOwner(owner framework.Updating) {
	fc.Collision.SetEntity(owner.(framework.ICollisionOwner))
	fc.Component.SetOwner(owner)
}

func (fc *FixedCollision) Update(dt float64) {
}
