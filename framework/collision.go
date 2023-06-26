package framework

import (
	"math"
)

type ISizable interface {
	GetSize() Size
}

type ICollisionOwner interface {
	IPhysicsObject
	GetScale() Vec2
	GetPivot() VecUV
}

type ICollisionSizableOwner interface {
	ISizable
	ICollisionOwner
}

type ICollisionComponentOwner interface {
	ICollisionSizableOwner
	Updating
}

type ICollisionFigure interface {
	Intersect(other ICollisionFigure) ContactSet
	Bounds() Bounds
	SetOffset(offset Vec2)
	SetScale(scale Vec2)
	SetRotation(rot Radian)
	GetOwner() ICollisionOwner
}

type Collide struct {
	Collision *Collision
	Contacts  []ContactSet
}

func NewCollide(collision *Collision, contacts []ContactSet) *Collide {
	return &Collide{
		Collision: collision,
		Contacts:  contacts,
	}
}

type Collision struct {
	Figures            []ICollisionFigure
	BehaviourOnCollide func(collide *Collide)
	entity             ICollisionOwner
	f                  *Framework
}

func InitCollision(figure ICollisionFigure) *Collision {
	c := &Collision{
		Figures: []ICollisionFigure{figure},
	}

	c.BehaviourOnCollide = c.OnCollide

	return c
}

func (c *Collision) GetEntity() ICollisionOwner {
	return c.entity
}

func (c *Collision) SetEntity(entity ICollisionOwner) {
	c.entity = entity
}

func NewPolygonCollision(points []Vec2, owner ICollisionOwner) ICollisionFigure {
	return &CollisionShapePolygon{
		points: points,
		owner:  owner,
	}
}

func NewPolygonCollisionUV(pointsUV []VecUV, size Size, owner ICollisionOwner) ICollisionFigure {
	points := make([]Vec2, len(pointsUV))
	for i, p := range pointsUV {
		points[i] = p.ToVec2(size)
	}
	p := &CollisionShapePolygon{
		points: points,
		owner:  owner,
	}
	p.SetOffset(Vec2{}.Sub(owner.GetPivot().ToVec2(size)))

	return p
}

func NewCircleCollision(size Size, owner ICollisionOwner) ICollisionFigure {
	return &CollisionShapeCircle{
		center: Vec2{size.Length / 2, size.Height / 2},
		radius: math.Min(size.Length, size.Height) / 2,
		owner:  owner,
	}
}

func NewBoxCollision(size Size, owner ICollisionOwner) ICollisionFigure {
	p1 := VecUV{0, 0}
	p2 := VecUV{1, 0}
	p3 := VecUV{1, 1}
	p4 := VecUV{0, 1}

	return NewPolygonCollisionUV([]VecUV{p1, p2, p3, p4}, size, owner)
}

func (c *Collision) AddFigure(f ICollisionFigure) {
	c.Figures = append(c.Figures, f)
}

func (c *Collision) GetFigures() []ICollisionFigure {
	return c.Figures
}

func (c *Collision) Intersect(collision *Collision) []ContactSet {
	res := make([]ContactSet, 0)
	for _, f := range c.Figures {
		for _, f2 := range collision.Figures {
			contactSet := f.Intersect(f2)
			if contactSet.WasIntersect() {
				res = append(res, contactSet)
			}
		}
	}
	return res
}

func (c *Collision) OnCollide(collide *Collide) {
	c.f.physic.ProcessingCollide(c.GetEntity(), collide)
}

func (c *Collision) Start(f *Framework) {
	c.f = f
	f.RegisterCollision(c, c.GetEntity())
}

func (c *Collision) Update(dt float64) {
	collides := make([]*Collide, 0)
	for _, collision := range c.f.GetClosestCollisonsFor(c) {
		cs := c.Intersect(collision)
		if len(cs) > 0 && cs[0].MoveOut != nil {
			collides = append(collides, NewCollide(collision, cs))
		}
	}
	for _, collide := range collides {
		func(collide *Collide) {
			c.f.AddAfterUpdate(func() {
				c.BehaviourOnCollide(collide)
			})
		}(collide)
	}
}

type ContactSet struct {
	Points  []Vec2
	MoveOut *Vec2
	Center  *Vec2
}

func (cs *ContactSet) WasIntersect() bool {
	return cs.Center != nil
}
