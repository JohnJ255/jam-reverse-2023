package game

import (
	"math"
	"reverse-jam-2023/components"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/models"
)

type Level4 struct {
}

func (l *Level4) Fill(level *Level) {
	car := models.NewSportCar(0)
	playerCar := entities.NewCar(framework.Player, car)
	playerCar.AddComponent(components.NewPlayerCarControl(level.size))
	playerCar.AddComponent(components.NewCarCollision(playerCar))
	level.player = playerCar
	level.AddEntity(playerCar)

	player := level.player

	player.Car.Position.X = 430
	player.Car.Position.Y = 350
	player.Car.Position.Angle = framework.Degrees(30).ToRadians()

	trailer1 := entities.NewTrailerToBackOfTractor(player.Car, player.Car.Size, 100, models.TrailerType(1))
	trailer1.AddComponent(components.NewTrailerCollision(trailer1))
	level.AddEntity(trailer1)
	player.Car.ConnectTrailer(trailer1.Trailer)

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
			if tr, ok := entity.(*entities.TrailerEntity); ok && math.Abs(float64(tr.GetRotation())) < math.Pi/4 {
				level.Change(level.framework, level.index+1)
			}
		})
	trigger.AddComponent(components.NewFixedCollision(trigger))
	level.AddEntity(trigger)
}
