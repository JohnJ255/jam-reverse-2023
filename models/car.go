package models

import (
	"fmt"
	"math"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/helper"
)

const massEtalon = 1000

type TrailerJoin interface {
	getSelfMass() float64
	getFullMass() float64
	getFrictionForce() float64
	GetPivot() helper.VecUV
	AddTraktor(towbar TowbarInterface)
	Control(velocity helper.Vec2)
}

type Car struct {
	Trailer       TrailerJoin
	Position      helper.DirectionPosition
	Size          helper.Size
	WheelBase     float64
	Pivot         helper.VecUV
	speed         float64
	powerful      float64
	minSpeed      float64
	maxSpeed      float64
	handling      float64 // управляемость
	speedHandling float64 // доля снижаемости управляемости на скорости
	wheelAngle    helper.Radian
	maxWheelAngle helper.Radian
	health        int
	maxHealth     int
	baseInertion  float64
	mass          float64
}

func NewSportCar(angle helper.Degrees) *Car {
	return &Car{
		Position: helper.DirectionPosition{
			Angle: angle.ToRadians(),
		},
		Size: helper.Size{
			Width:  56,
			Length: 114,
		},
		Pivot:         helper.VecUV{0.2, 0.5},
		powerful:      160,
		minSpeed:      helper.KmphToPixelsPerTick(-54),
		maxSpeed:      helper.KmphToPixelsPerTick(180),
		handling:      0.02,
		speedHandling: 0.7,
		maxWheelAngle: helper.Degrees(45).ToRadians(),
		health:        100,
		maxHealth:     100,
		WheelBase:     80,
		baseInertion:  0.97,
		mass:          800,
	}
}

func (c *Car) Control(accelerate float64, wheelRotation float64) {
	powerful := c.powerful / c.mass
	k := 1 + (massEtalon-c.mass)/massEtalon
	minSpeed := c.minSpeed * k
	maxSpeed := c.maxSpeed * k
	inertion := c.calcInertionDependsMass()
	if accelerate == 0 && c.speed != 0 {
		c.speed *= inertion
		if math.Abs(c.speed) < powerful {
			c.speed = 0
		}
	} else {
		c.speed = helper.Limited(c.speed+accelerate*powerful, minSpeed, maxSpeed)
	}
	maxWheelAngle := helper.Radian(float64(c.maxWheelAngle) * (maxSpeed - math.Abs(c.speed)*c.speedHandling) / maxSpeed)
	newWheelAngle := helper.Radian(helper.Stepped(float64(c.wheelAngle), wheelRotation*float64(maxWheelAngle), c.handling))
	c.wheelAngle = helper.Limited(newWheelAngle, -maxWheelAngle, maxWheelAngle)

	c.Position.Angle += helper.Radian(math.Atan2(c.WheelBase*math.Tan(float64(c.wheelAngle)), c.Size.Length+c.WheelBase) * c.speed * 0.03)
	dx := c.speed * math.Cos(float64(c.Position.Angle))
	dy := c.speed * math.Sin(float64(c.Position.Angle))
	c.Position.X += dx
	c.Position.Y += dy

	if c.Trailer != nil {
		c.Trailer.Control(helper.Vec2{dx, dy})
	}

	framework.DebugWatchAdd("Speed", func() string {
		return fmt.Sprintf("%f", c.speed)
	})
	framework.DebugWatchAdd("newWheelAngle", func() string {
		return fmt.Sprintf("%f", newWheelAngle)
	})
	framework.DebugWatchAdd("wheelAngle", func() string {
		return fmt.Sprintf("%f", c.wheelAngle)
	})
	framework.DebugWatchAdd("maxWheelAngle", func() string {
		return fmt.Sprintf("%f", maxWheelAngle)
	})
}

func (c *Car) getSelfMass() float64 {
	return c.mass
}

func (c *Car) getFullMass() float64 {
	if c.Trailer != nil {
		return c.mass + c.Trailer.getFullMass()
	}

	return c.getSelfMass()
}

func (c *Car) getFrictionForce() float64 {
	return 1 - c.calcInertionDependsMass()
}

func (c *Car) AddTrailer(cargo TrailerJoin) {
	c.Trailer = cargo
	cargo.AddTraktor(c)
}

func (c *Car) calcInertionDependsMass() float64 {
	mass := c.mass
	if c.Trailer != nil {
		mass += c.Trailer.getSelfMass()
	}
	k := 1 + (massEtalon-mass)/massEtalon
	return helper.Limited(c.baseInertion-k/10, 0.9, 0.999)
}

func (c *Car) GetPivot() helper.VecUV {
	return c.Pivot
}

func (c *Car) GetPosition() helper.DirectionPosition {
	return c.Position
}

func (c *Car) GetTowbarPosition() helper.Vec2 {
	x := c.Position.X - c.Size.Length*c.Pivot.U*math.Cos(float64(c.Position.Angle))
	y := c.Position.Y - c.Size.Width*c.Pivot.V*math.Sin(float64(c.Position.Angle))*0.8
	return helper.Vec2{x, y}
}
