package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/helper"
	"reverse-jam-2023/models"
	"strconv"
)

type Game struct {
	player     *entities.CarEntity
	level      framework.Drawing
	WindowSize helper.IntSize
	scale      float64
}

func NewGame() *Game {
	car := models.NewSportCar(0)
	return &Game{
		player: entities.NewCar(framework.Player, car),
		level:  entities.NewLevel(0),
		WindowSize: helper.IntSize{
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
		trailer := models.NewTrailerToBackOfTractor(p.Car, p.Car.Size, 400, models.TrailerType(trType))
		p.Car.AddTrailer(trailer)
		return "trailer added"
	})
	f.SetConsoleCommand("towbar", func(params ...string) string {
		p := g.player
		if params[0] == "1" {
			f.SetDebugDraw("towbar", func(screen *ebiten.Image) {
				pos := p.Car.GetTowbarPosition()
				vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 4, color.NRGBA{0, 255, 0, 255}, false)
				pos = p.Car.GetPosition().Position
				vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 4, color.NRGBA{0, 0, 255, 255}, false)
			})
			return "towbar added"
		}
		f.RemoveDebugDraw("towbar")
		return "towbar removed"
	})
	f.MakeConsoleCommand("towbar 1")
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
