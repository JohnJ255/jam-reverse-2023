package game

import (
	"math"
	"reverse-jam-2023/components"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/models"
)

type Level6 struct {
}

func (l *Level6) GetSize() framework.Size {
	return framework.Size{800, 800}
}

func (l *Level6) Fill(level *LevelManager) {
	l1 := &Level1{}
	l3 := &Level3{}
	l1.makePlayerCar(level)

	player := level.player

	player.Car.Position.X = 530
	player.Car.Position.Y = 250
	player.Car.Position.Angle = framework.Degrees(180).ToRadians()

	trailer1 := entities.NewTrailerToBackOfTractor(player.Car, player.Car.Size, 100, models.TrailerType(1))
	trailer1.AddComponent(components.NewTrailerCollision(trailer1))
	level.AddEntity(trailer1)
	player.Car.ConnectTrailer(trailer1.Trailer)

	l3.makeParkingCar(level, 312, 320, math.Pi/2-0.2)
	l1.makeParkingCar(level, 212, 380)
	l1.makeParkingCar(level, 148, 380)
	l1.makeParkingCar(level, 520, 380)
	l3.makeParkingCar(level, 520, 630, -math.Pi/2)
	l3.makeParkingCar(level, 650, 640, -math.Pi/2)

	l1.makeWallsCollisions(level)
	l3.makeGoalTrigger(level)
}
