package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/components"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/loader"
	"reverse-jam-2023/models"
	"strconv"
)

type Level struct {
	*framework.Sprite
	name   string
	size   framework.Size
	player *entities.CarEntity
	camera framework.ICamera
}

func NewLevel(index int) *Level {
	bgSize := framework.Size{1200, 600}
	level := &Level{
		Sprite: framework.InitSprites(bgSize),
		name:   "level " + strconv.Itoa(index),
		size:   bgSize,
	}
	level.camera = framework.NewStaticCamera(level, level.Sprite)
	level.LoadResources(&loader.ResourceLoader{}, loader.LevelFileNames[index])
	return level
}

func (l *Level) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	return l.camera.GetTransforms(scale)
}

func (l *Level) Init(f *framework.Framework) {
	car := models.NewSportCar(0)
	playerCar := entities.NewCar(framework.Player, car)
	playerCar.AddComponent(components.NewPlayerCarControl())
	playerCar.AddComponent(components.NewCarCollision(playerCar))
	l.player = playerCar
	f.AddEntity(playerCar)

	if l.name == "level 1" {
		car.Position.X = 200
		car.Position.Y = 300
		car.Position.Angle = framework.Degrees(-45).ToRadians()

		trailer1 := entities.NewTrailer(framework.NewDPos(300, 100, framework.Degrees(45).ToRadians()), car.Size, 100, models.TrailerType(1))
		trailer1.AddComponent(components.NewTrailerCollision(trailer1))
		f.AddEntity(trailer1)

		xmcar := models.NewSportCar(framework.Degrees(90))
		xmcar.Position.X = 400
		xmcar.Position.Y = 300
		xcar := entities.NewCar(framework.Computer, xmcar)
		xcar.AddComponent(components.NewCarCollision(xcar))
		f.AddEntity(xcar)
	}

	if l.name == "level 2" {
		car.Position.X = 200
		car.Position.Y = 300
		car.Position.Angle = framework.Degrees(25).ToRadians()

		trailer1 := entities.NewTrailerToBackOfTractor(car, car.Size, 100, models.TrailerType(1))
		trailer1.AddComponent(components.NewTrailerCollision(trailer1))
		f.AddEntity(trailer1)

		car.ConnectTrailer(trailer1.Trailer)

		xmcar := models.NewSportCar(framework.Degrees(90))
		xmcar.Position.X = 400
		xmcar.Position.Y = 300
		xcar := entities.NewCar(framework.Computer, xmcar)
		xcar.AddComponent(components.NewCarCollision(xcar))
		f.AddEntity(xcar)
	}
}

func (l *Level) GetPlayer() framework.Entity {
	return l.player
}

func (l *Level) Update(dt float64) {
	l.camera.Control(l.player)
}

func (l *Level) GetSize() framework.Size {
	return l.size
}
