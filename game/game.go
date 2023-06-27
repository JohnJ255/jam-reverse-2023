package game

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"image/color"
	"log"
	"reverse-jam-2023/components"
	"reverse-jam-2023/entities"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/models"
	"strconv"
)

type Game struct {
	level      *Level
	WindowSize framework.IntSize
	scale      float64
	f          *framework.Framework
	fontGUI    font.Face
}

func NewGame() *Game {
	ttf, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}
	faceOpt := &truetype.Options{
		Size:    14,
		DPI:     72,
		Hinting: font.HintingFull,
	}

	g := &Game{
		WindowSize: framework.IntSize{
			Width:  800,
			Height: 600,
		},
		scale:   0.1,
		fontGUI: truetype.NewFace(ttf, faceOpt),
	}
	g.level = NewLevel(1, g)

	return g
}

func (g *Game) Start(f *framework.Framework) {
	f.DebugModeEnable()
	g.level.Init(f)
	g.f = f

	f.SetConsoleCommand("trailer", func(params ...string) string {
		p := g.level.GetPlayer()
		trType, err := strconv.Atoi(params[0])
		if err != nil {
			f.MessageToConsole("invalid parameter: need trailer type")
		}
		var trailer *entities.TrailerEntity
		car := p.GetModel().(*models.Car)
		if trType == 1 {
			trailer = entities.NewTrailerToBackOfTractor(car, car.Size, 400, models.TrailerType(trType))
			car.ConnectTrailer(trailer.Trailer)
		} else {
			trailer = entities.NewTrailer(framework.NewDPos(200, 200, 0), car.Size, 400, models.TrailerType(trType))
		}
		if trailer != nil {
			trailer.AddComponent(components.NewTrailerCollision(trailer))
			f.AddEntity(trailer)
		}
		return "trailer added"
	})
	f.SetConsoleCommand("towbar", func(params ...string) string {
		car := g.level.GetPlayer().GetModel().(*models.Car)
		if params[0] == "1" {
			f.Debug.SetDebugDraw("towbar", func(screen *ebiten.Image) {
				pos := car.GetTowbarPosition()
				vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 4, color.NRGBA{0, 255, 0, 255}, false)
				pos = car.GetPosition().Vec2
				vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 4, color.NRGBA{0, 0, 255, 255}, false)
				if car.Trailer != nil {
					pos = car.Trailer.(*models.Trailer).Position.Vec2
					vector.DrawFilledCircle(screen, float32(pos.X), float32(pos.Y), 3, color.NRGBA{160, 160, 65, 255}, false)
					pos = pos.Add(car.Trailer.(*models.Trailer).GetTowbarLocalPosition())
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

	//f.MakeConsoleCommand("towbar 1")
	//f.MakeConsoleCommand("trailer 1")
	//f.MakeConsoleCommand("show_collisions 1")
}

func (g *Game) Update(dt float64) error {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		g.level.Change(g.f, g.level.index+1)
	}

	g.level.Update(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.level != nil {
		screen.DrawImage(g.level.GetSprite(), g.SceneTransform(g.level.GetTransforms(1)))
	}
}

func (g *Game) DrawGUI(screen *ebiten.Image) {
	text.Draw(screen, g.level.name, g.fontGUI, 300, 15, color.White)
	text.Draw(screen, fmt.Sprintf("Score: %d", g.level.Score), g.fontGUI, 400, 15, color.White)
}

func (g *Game) GetTitle() string {
	return "reverse-jam-2023"
}

func (g *Game) SceneTransform(transforms *ebiten.DrawImageOptions) *ebiten.DrawImageOptions {
	pos := g.level.camera.GetPosition()

	transforms.GeoM.Translate(-pos.X, -pos.Y)

	return transforms
}
