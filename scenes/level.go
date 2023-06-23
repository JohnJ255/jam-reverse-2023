package scenes

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
	player framework.Entity
}

func NewLevel(index int) *Level {
	level := &Level{
		Sprite: framework.InitSprites(),
		name:   "level " + strconv.Itoa(index),
		size:   framework.Size{1200, 600},
	}
	level.LoadResources(&loader.ResourceLoader{}, loader.LevelFileNames[index])
	return level
}

func (l *Level) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	return l.PivotTransform(scale, l.size, framework.VecUV{})
}

func (l *Level) Init(index int, f *framework.Framework) {
	car := models.NewSportCar(0)
	playerCar := entities.NewCar(framework.Player, car)
	playerCar.AddComponent(components.NewPlayerCarControl())
	playerCar.AddComponent(components.NewCarCollision(playerCar))
	l.player = playerCar
	f.AddEntity(playerCar)

	if index == 1 {
		car.Position.X = 200
		car.Position.Y = 300
		car.Position.Angle = framework.Degrees(-45).ToRadians()

		trailer1 := entities.NewTrailer(framework.NewDPos(300, 100, framework.Degrees(45).ToRadians()), car.Size, 400, models.TrailerType(1))
		trailer1.AddComponent(components.NewTrailerCollision(trailer1))
		f.AddEntity(trailer1)
	}

	if index == 2 {
		car.Position.X = 200
		car.Position.Y = 300
		car.Position.Angle = framework.Degrees(25).ToRadians()

		trailer1 := entities.NewTrailerToBackOfTractor(car, car.Size, 400, models.TrailerType(1))
		trailer1.AddComponent(components.NewTrailerCollision(trailer1))
		f.AddEntity(trailer1)

		car.ConnectTrailer(trailer1.Trailer)
	}
}

func (l *Level) GetPlayer() framework.Entity {
	return l.player
}
