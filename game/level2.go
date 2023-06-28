package game

import (
	"math"
	"reverse-jam-2023/framework"
)

type Level2 struct {
}

func (l *Level2) GetSize() framework.Size {
	return framework.Size{800, 800}
}

func (l *Level2) Fill(level *LevelManager) {
	l1 := &Level1{}
	l1.makeWallsCollisions(level)

	l1.makePlayerCar(level)

	player := level.player

	player.Car.Position.X = 330
	player.Car.Position.Y = 365
	player.Car.Position.Angle = math.Pi / 2

	l1.makeParkingCar(level, 212, 380)
	l1.makeParkingCar(level, 520, 380)
	l1.makeGoalTrigger(level)
}
