package game

import (
	"math"
	"reverse-jam-2023/components"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/models"
)

type Level1 struct {
}

const WallWidth = 15

func (l *Level1) Fill(level *Level) {
	car := models.NewSportCar(0)
	playerCar := entities.NewCar(framework.Player, car)
	playerCar.AddComponent(components.NewPlayerCarControl(level.size))
	playerCar.AddComponent(components.NewCarCollision(playerCar))
	level.player = playerCar
	level.AddEntity(playerCar)

	player := level.player

	player.Car.Position.X = 330
	player.Car.Position.Y = 250

	xmcar := models.NewSportCar(framework.Degrees(90))
	xmcar.Position.X = 200
	xmcar.Position.Y = 330
	xcar := entities.NewCar(framework.Computer, xmcar)
	xcar.AddComponent(components.NewCarCollision(xcar))
	level.AddEntity(xcar)

	wall := entities.NewWall(framework.Vec2{114, 200}, framework.Size{128, WallWidth})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{114, 285}, framework.Size{128, WallWidth})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{114, 200}, framework.Size{WallWidth, 85})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)

	wall = entities.NewWall(framework.Vec2{50, 75}, framework.Size{700, WallWidth})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{50, 705}, framework.Size{700, WallWidth})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{45, 75}, framework.Size{WallWidth, 705})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{750, 75}, framework.Size{WallWidth, 705})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)

	trigger := entities.NewTrigger(framework.Vec2{125, 220}, framework.Size{15, 60},
		func(entity framework.ICollisionOwner, collide *framework.Collide) {
			if car, ok := entity.(*entities.CarEntity); ok && math.Abs(float64(car.GetRotation())) < math.Pi/4 {
				level.Change(level.framework, level.index+1)
			}
		})
	trigger.AddComponent(components.NewFixedCollision(trigger))
	level.AddEntity(trigger)
}
