package game

import (
	"math"
	"reverse-jam-2023/components"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/models"
)

type Level3 struct {
}

func (l *Level3) GetSize() framework.Size {
	return framework.Size{800, 800}
}

func (l *Level3) Fill(level *LevelManager) {
	l1 := &Level1{}
	l1.makePlayerCar(level)

	player := level.player

	player.Car.Position.X = 430
	player.Car.Position.Y = 270

	trailer1 := entities.NewTrailerToBackOfTractor(player.Car, player.Car.Size, 100, models.TrailerType(1))
	trailer1.AddComponent(components.NewTrailerCollision(trailer1))
	level.AddEntity(trailer1)
	player.Car.ConnectTrailer(trailer1.Trailer)

	l1.makeParkingCar(level, 212, 380)
	l1.makeParkingCar(level, 520, 380)
	l.makeParkingCar(level, 520, 470, -math.Pi/2+0.07)

	l1.makeWallsCollisions(level)
	l.makeGoalTrigger(level)
}

func (l *Level3) makeGoalTrigger(level *LevelManager) {
	trigger := entities.NewTrigger(framework.Vec2{135, 235}, framework.Size{10, 60},
		func(entity framework.ICollisionOwner, collide *framework.Collide) {
			if tr, ok := entity.(*entities.TrailerEntity); ok && math.Abs(float64(tr.GetRotation().NormalizePi2())) < math.Pi/18 {
				level.Change(level.framework, level.index+1)
			}
		})
	trigger.AddComponent(components.NewFixedCollision(trigger))
	level.AddEntity(trigger)
}

func (l *Level3) makeParkingCar(level *LevelManager, x, y float64, rot framework.Radian) {
	l1 := &Level1{}
	c := l1.makeParkingCar(level, x, y)
	c.SetRotation(rot)
}
