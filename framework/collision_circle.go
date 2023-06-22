package framework

import "math"

type CollisionShapeCircle struct {
	center Vec2
	radius float64
	owner  ICollisionOwner
	//radiusX float64
	//radiusY float64
	//rotation Radian
}

func (c *CollisionShapeCircle) Intersect(other ICollisionFigure) ContactSet {
	switch other.(type) {
	case *CollisionShapeCircle:
		circle := other.(*CollisionShapeCircle)
		points := c.IntersectionPointsCircle(circle)
		if len(points) == 0 {
			return ContactSet{}
		}
		mo := Vec2{c.center.X - circle.center.X, c.center.Y - circle.center.Y}
		mo = mo.Normalize().Mul(c.radius + circle.radius - mo.Length())
		return ContactSet{
			Points:  points,
			MoveOut: &mo,
			Center:  CalcCenter(points),
		}
	case *CollisionShapePolygon:
		polygon := other.(*CollisionShapePolygon)
		contactSet := polygon.Intersect(c)
		if contactSet.WasIntersect() {
			mo := contactSet.MoveOut.Mul(-1)
			contactSet.MoveOut = &mo
		}
		return contactSet
	default:
		return ContactSet{}
	}
}

func (c *CollisionShapeCircle) SetOffset(offset Vec2) {
	c.center.X += offset.X
	c.center.Y += offset.Y
}

func (c *CollisionShapeCircle) SetScale(scale Vec2) {
	//c.radiusX *= scale.X
	//c.radiusY *= scale.Y
	c.radius *= scale.X
}

func (c *CollisionShapeCircle) SetRotation(rot Radian) {
	//c.rotation = rot
}

func (c *CollisionShapeCircle) Bounds() Bounds {
	return Bounds{
		//Min: Vec2{c.center.X - c.radiusX, c.center.Y - c.radiusY},
		//Max: Vec2{c.center.X + c.radiusX, c.center.Y + c.radiusY},
		Min: Vec2{c.center.X - c.radius, c.center.Y - c.radius},
		Max: Vec2{c.center.X + c.radius, c.center.Y + c.radius},
	}
}

func (c *CollisionShapeCircle) IntersectionPointsCircle(circle *CollisionShapeCircle) []Vec2 {
	dist := math.Sqrt(math.Pow(circle.center.X-c.center.X, 2) + math.Pow(circle.center.Y-c.center.Y, 2))
	if dist > (c.radius+circle.radius) || dist < math.Abs(c.radius-circle.radius) || dist == 0 && c.radius == circle.radius {
		return nil
	}

	a := (math.Pow(c.radius, 2) - math.Pow(circle.radius, 2) + math.Pow(dist, 2)) / (2 * dist)
	h := math.Sqrt(math.Pow(c.radius, 2) - math.Pow(a, 2))

	x := c.center.X + a*(circle.center.X-c.center.X)/dist
	y := c.center.Y + a*(circle.center.Y-c.center.Y)/dist

	x2 := h * (circle.center.Y - c.center.Y) / dist
	y2 := h * (circle.center.X - c.center.X) / dist

	return []Vec2{
		{x + x2, y - y2},
		{x - x2, y + y2},
	}
}

func (c *CollisionShapeCircle) GetOwner() ICollisionOwner {
	return c.owner
}

func (c *CollisionShapeCircle) GetCenter() Vec2 {
	return c.center.Add(c.owner.GetPosition())
}

func (c *CollisionShapeCircle) GetRadius() float64 {
	return c.radius * math.Min(c.owner.GetScale().X, c.owner.GetScale().Y)
}
