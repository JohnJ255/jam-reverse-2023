package framework

type IPhysicsObject interface {
	IPositioning
	IsFixed() bool
	GetMass() float64
	GetFriction() float64
}

type PhysicTop struct {
}

func (p *PhysicTop) ProcessingCollide(obj IPhysicsObject, collide *Collide) {
	if collide.Collision.GetEntity() == nil || len(collide.Contacts) == 0 {
		return
	}
	if other, ok := collide.Collision.GetEntity().(IPhysicsObject); ok {
		pos := obj.GetPosition()
		for _, cs := range collide.Contacts {
			pos.X += cs.MoveOut.X
			pos.Y += cs.MoveOut.Y

			sign := p.calcAngleStep(other.GetMass() / obj.GetMass())
			if obj.GetRotation().LefterThan((*cs.MoveOut).ToRadian()) {
				sign *= -1
			}
			obj.SetRotation(obj.GetRotation() + sign)

		}
		obj.SetPosition(pos)
	}
}

func (p *PhysicTop) calcAngleStep(f float64) Radian {
	return Radian(0.01 * f)
}
