package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"reverse-jam-2023/components"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/models"
	"strconv"
)

type Game struct {
	player     *entities.CarEntity
	level      framework.Drawing
	WindowSize framework.IntSize
	scale      float64
}

func NewGame() *Game {
	car := models.NewSportCar(0)
	car.Position.X = 300
	car.Position.Y = 100
	playerCar := entities.NewCar(framework.Player, car)
	playerCar.AddComponent(components.NewPlayerCarControl())
	playerCar.AddComponent(components.NewCarCollision(playerCar))
	return &Game{
		player: playerCar,
		level:  entities.NewLevel(1),
		WindowSize: framework.IntSize{
			Width:  800,
			Height: 600,
		},
		scale: 0.1,
	}
}

func (g *Game) Start(f *framework.Framework) {
	f.AddEntity(g.player)
	f.DebugModeEnable()
	f.SetConsoleCommand("trailer", func(params ...string) string {
		p := g.player
		trType, err := strconv.Atoi(params[0])
		if err != nil {
			f.MessageToConsole("invalid parameter: need trailer type")
		}
		var trailer *entities.TrailerEntity
		if trType == 1 {
			trailer = entities.NewTrailerToBackOfTractor(p.Car, p.Car.Size, 400, models.TrailerType(trType))
			p.Car.AddTrailer(trailer.Trailer)
		} else {
			trailer = entities.NewTrailer(framework.NewDPos(200, 200, 0), p.Car.Size, 400, models.TrailerType(trType))
		}
		if trailer != nil {
			trailer.AddComponent(components.NewTrailerCollision(trailer, g.player.GetComponent("CarCollision")))
			f.AddEntity(trailer)
		}
		return "trailer added"
	})
	f.SetConsoleCommand("towbar", func(params ...string) string {
		p := g.player
		if params[0] == "1" {
			f.Debug.SetDebugDraw("towbar", func(screen *ebiten.Image) {
				pos := p.Car.GetTowbarPosition()
				vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 4, color.NRGBA{0, 255, 0, 255}, false)
				pos = p.Car.GetPosition().Vec2
				vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 4, color.NRGBA{0, 0, 255, 255}, false)
				if p.Car.Trailer != nil {
					pos = p.Car.Trailer.(*models.Trailer).Position.Vec2
					vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 3, color.NRGBA{160, 160, 65, 255}, false)
					pos = pos.Add(p.Car.Trailer.(*models.Trailer).GetTowbarLocalPosition())
					vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 3, color.NRGBA{160, 0, 65, 255}, false)
				}
			})
			return "towbar added"
		}
		f.Debug.RemoveDebugDraw("towbar")
		return "towbar removed"
	})
	f.SetConsoleCommand("show_collisions", func(params ...string) string {
		if params[0] == "1" {
			f.Debug.SetDebugDraw("collisions", f.Debug.DefaultDrawCollisions)
			return "collisions debug drawing enable"
		}
		f.Debug.RemoveDebugDraw("collisions")
		return "collisions debug drawing disabled"
	})

	f.MakeConsoleCommand("towbar 1")
	f.MakeConsoleCommand("trailer 1")
	f.MakeConsoleCommand("show_collisions 1")
}

func (g *Game) Update(dt float64) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.level != nil {
		screen.DrawImage(g.level.GetSprite(), g.level.GetTransforms(1))
	}
}

func (g *Game) GetTitle() string {
	return "reverse-jam-2023"
}
