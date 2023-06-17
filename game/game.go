package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/helper"
	"reverse-jam-2023/models"
)

type Game struct {
	player     framework.Controlling
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
	f.DebugModeEnable()
}

func (g *Game) Update(dt float64) error {
	accelerate := 0.0
	wheel := 0.0
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		wheel = -1
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		wheel = 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		accelerate = 1.0
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		accelerate = -0.3
	}
	g.player.Control(map[string]float64{
		"accelerate": accelerate,
		"wheel":      wheel,
	})

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.level.GetSprite(), g.level.GetTransforms(1))

	screen.DrawImage(g.player.GetSprite(), g.player.GetTransforms(g.scale))
}

func (g *Game) GetTitle() string {
	return "reverse-jam-2023"
}
