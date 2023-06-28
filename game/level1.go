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

const WallWidth = 25

func (l *Level1) GetSize() framework.Size {
	return framework.Size{800, 800}
}

func (l *Level1) Fill(level *LevelManager) {
	l.makePlayerCar(level)

	player := level.player
	player.Car.Position.X = 330
	player.Car.Position.Y = 265

	l.makeParkingCar(level, 212, 380)

	l.makeWallsCollisions(level)

	l.makeGoalTrigger(level)
}

func (l *Level1) makeWallsCollisions(level *LevelManager) {
	wall := entities.NewWall(framework.Vec2{114, 200}, framework.Size{130, WallWidth})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{114, 300}, framework.Size{130, WallWidth})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{114, 225}, framework.Size{WallWidth, 75})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{40, 300}, framework.Size{75, WallWidth})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)

	wall = entities.NewWall(framework.Vec2{40, 70}, framework.Size{720, WallWidth})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{50, 705}, framework.Size{720, WallWidth})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{40, 70}, framework.Size{WallWidth, 705})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
	wall = entities.NewWall(framework.Vec2{745, 70}, framework.Size{WallWidth, 705})
	wall.AddComponent(components.NewFixedCollision(wall))
	level.AddEntity(wall)
}

func (l *Level1) makeGoalTrigger(level *LevelManager) {
	trigger := entities.NewTrigger(framework.Vec2{135, 235}, framework.Size{10, 60},
		func(entity framework.ICollisionOwner, collide *framework.Collide) {
			if car, ok := entity.(*entities.CarEntity); ok && math.Abs(float64(car.GetRotation().NormalizePi2())) < math.Pi/18 {
				level.Change(level.framework, level.index+1)
			}
		})
	trigger.AddComponent(components.NewFixedCollision(trigger))
	level.AddEntity(trigger)
}

func (l *Level1) makeParkingCar(level *LevelManager, x, y float64) *entities.CarEntity {
	xmcar := models.NewSportCar(framework.Degrees(90))
	xmcar.Position.X = x
	xmcar.Position.Y = y
	xcar := entities.NewCar(framework.Computer, xmcar)
	xcar.AddComponent(components.NewCarCollision(xcar))
	level.AddEntity(xcar)
	return xcar
}

func (l *Level1) makePlayerCar(level *LevelManager) {
	car := models.NewSportCar(0)
	playerCar := entities.NewCar(framework.Player, car)
	playerCar.AddComponent(components.NewPlayerCarControl(level.size, level))
	playerCar.AddComponent(components.NewCarCollision(playerCar))
	level.player = playerCar
	level.AddEntity(playerCar)
}
