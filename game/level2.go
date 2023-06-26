package game

import (
	"reverse-jam-2023/components"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/models"
)

type Level2 struct {
}

func (l *Level2) Fill(level *Level) {
	car := models.NewSportCar(0)
	playerCar := entities.NewCar(framework.Player, car)
	playerCar.AddComponent(components.NewPlayerCarControl(level.size))
	playerCar.AddComponent(components.NewCarCollision(playerCar))
	level.player = playerCar
	level.AddEntity(playerCar)

	player := level.player

	player.Car.Position.X = 200
	player.Car.Position.Y = 300
	player.Car.Position.Angle = framework.Degrees(25).ToRadians()

	trailer1 := entities.NewTrailerToBackOfTractor(player.Car, player.Car.Size, 100, models.TrailerType(1))
	trailer1.AddComponent(components.NewTrailerCollision(trailer1))
	level.AddEntity(trailer1)

	player.Car.ConnectTrailer(trailer1.Trailer)

	xmcar := models.NewSportCar(framework.Degrees(90))
	xmcar.Position.X = 400
	xmcar.Position.Y = 300
	xcar := entities.NewCar(framework.Computer, xmcar)
	xcar.AddComponent(components.NewCarCollision(xcar))
	level.AddEntity(xcar)
}
