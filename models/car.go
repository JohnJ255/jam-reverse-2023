package models

import (
	"fmt"
	"math"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/helper"
)

type Car struct {
	Position      helper.DirectionPosition
	Size          helper.Size
	WheelBase     float64
	speed         float64
	powerful      float64
	minSpeed      float64
	maxSpeed      float64
	handling      float64 // управляемость
	wheelAngle    float64
	maxWheelAngle float64
	health        int
	maxHealth     int
	inertion      float64
	mass          float64
}

const massEtalon = 1000

func NewSportCar(angle helper.Degrees) *Car {
	return &Car{
		Position: helper.DirectionPosition{
			Angle: helper.ToRadians(angle),
		},
		Size: helper.Size{
			Width:  40,
			Length: 100,
		},
		powerful:      160,
		minSpeed:      helper.KmphToPixelsPerTick(-54),
		maxSpeed:      helper.KmphToPixelsPerTick(180),
		handling:      0.5,
		maxWheelAngle: helper.ToRadians(45),
		health:        100,
		maxHealth:     100,
		WheelBase:     80,
		inertion:      0.95,
		mass:          800,
	}
}

func (c *Car) Control(accelerate float64, wheelRotation float64) {
	powerful := c.powerful / c.mass
	k := 1 + (massEtalon-c.mass)/massEtalon
	minSpeed := c.minSpeed * k
	maxSpeed := c.maxSpeed * k
	if accelerate == 0 && c.speed != 0 {
		c.speed *= c.inertion
		if math.Abs(c.speed) < powerful {
			c.speed = 0
		}
	} else {
		c.speed = helper.Limited(c.speed+accelerate*powerful, minSpeed, maxSpeed)
	}
	c.wheelAngle = helper.Limited((c.wheelAngle+wheelRotation)*c.handling, -c.maxWheelAngle, c.maxWheelAngle)
	if math.Abs(c.speed) > 2 {
		c.wheelAngle *= (maxSpeed - math.Abs(c.speed)/2) / maxSpeed
	}
	c.Position.Angle += math.Atan2(c.WheelBase*math.Tan(c.wheelAngle), c.Size.Length+c.WheelBase) * c.speed * 0.03
	c.Position.X += c.speed * math.Cos(c.Position.Angle)
	c.Position.Y += c.speed * math.Sin(c.Position.Angle)

	//fmt.Printf("speed: %f, wheelAngle: %f, angle: %f, x: %f, y: %f\n", c.speed, c.wheelAngle, c.Position.Angle, c.Position.X, c.Position.Y)
	framework.DebugWatchAdd("Speed", func() string {
		return fmt.Sprintf("%f", c.speed)
	})

}
