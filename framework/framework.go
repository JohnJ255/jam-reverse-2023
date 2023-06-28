package framework

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"strings"
	"time"
)

type Game interface {
	Start(f *Framework)
	Update(dt float64) error
	Draw(screen *ebiten.Image)
	SceneTransform(transforms *ebiten.DrawImageOptions) *ebiten.DrawImageOptions
	DrawGUI(screen *ebiten.Image)
	IsPaused() bool
}

type IPhysicsEngine interface {
	ProcessingCollide(obj IPhysicsObject, collide *Collide)
}

type Framework struct {
	game         Game
	entities     []Entity
	collisions   []*Collision
	lastUpdate   time.Time
	console      *Console
	windowWidth  int
	windowHeight int
	windowTitle  string
	ticks        uint64
	Debug        *DebugTool
	WorldStarted bool
	afterUpdates []func()
	physic       IPhysicsEngine
	Events       *EventSystem
	Audio        *AudioPlayer
}

var fw *Framework

func InitWindowGame(g Game, windowWidth, windowHeight int, windowTitle string, ttf *truetype.Font) *Framework {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle(windowTitle)
	fw = &Framework{
		game:         g,
		entities:     make([]Entity, 0),
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		windowTitle:  windowTitle,
		console:      NewConsole(ttf),
		collisions:   make([]*Collision, 0),
		physic:       &PhysicTop{},
		Events:       NewEventSystem(),
	}
	fw.Debug = NewDebugTool(fw, &DefaultCollisionPainter{
		color: color.NRGBA{40, 255, 40, 255},
	}, &DefaultIntersectionPainter{
		color:       color.NRGBA{255, 155, 155, 255},
		arrowColor:  color.NRGBA{255, 255, 255, 255},
		arrowLength: 10,
	})
	return fw
}

func (f *Framework) Run() error {
	return ebiten.RunGame(f)
}

func (f *Framework) Update() error {
	now := time.Now()
	dt := now.Sub(f.lastUpdate).Seconds()
	f.lastUpdate = now
	f.ticks++
	entities := f.entities
	if f.ticks == 1 {
		f.game.Start(f)

		entities = f.entities
		for _, e := range entities {
			e.Start(f)
		}
		f.WorldStarted = true
		return nil
	}

	if f.console.IsAvailable && inpututil.IsKeyJustPressed(f.console.ToggleKey) {
		f.console.Toggle()
	}

	if f.console.IsOpened {
		f.console.Update(f)

		return nil
	}

	err := f.game.Update(dt)

	if f.game.IsPaused() {
		return err
	}

	for _, e := range entities {
		e.Update(dt)
	}

	for _, afterUpdate := range f.afterUpdates {
		afterUpdate()
	}
	f.afterUpdates = make([]func(), 0, len(f.afterUpdates))

	f.Audio.Update()

	return err
}

func (f *Framework) Draw(screen *ebiten.Image) {
	f.game.Draw(screen)
	for _, e := range f.entities {
		spr := e.GetSprite()
		if spr != nil {
			screen.DrawImage(spr, f.game.SceneTransform(e.GetTransforms(1)))
		}
	}

	f.game.DrawGUI(screen)

	for _, drawer := range f.Debug.Draws {
		drawer(screen)
	}
	if f.console.IsOpened {
		f.console.Draw(screen, 0, 0, f.windowWidth, f.windowHeight/3)
	}
}

func (f *Framework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func (f *Framework) SetConsoleCommand(name string, do func(params ...string) string) {
	f.console.SetCommand(name, do)
}

func DebugWatchAdd(name string, valSource func() string) {
	if fw == nil {
		fmt.Println("DebugWatchAdd before init framework!")
		return
	}
	fw.console.AddWatch(name, valSource)
}

func (f *Framework) DebugModeEnable() {
	f.console.IsAvailable = true
	f.console.Println("Debug mode enabled")
}

func (f *Framework) MessageToConsole(msg string) {
	f.console.Println(msg)
}

func (f *Framework) MakeConsoleCommand(s string) {
	f.console.Println(s)
	params := strings.Split(s, " ")
	if len(params) == 1 {
		f.console.makeCommand(params[0])
	} else {
		f.console.makeCommand(params[0], params[1:]...)
	}
}

func (f *Framework) AddEntity(entity Entity) {
	f.entities = append(f.entities, entity)
	if f.WorldStarted {
		entity.Start(f)
	}
}

func (f *Framework) RemoveEntity(entity Entity) {
	res := make([]Entity, 0, len(f.entities))
	for _, ent := range f.entities {
		if ent != entity {
			res = append(res, ent)
		}
	}
	f.entities = res
}

// TODO: refactoring to find only closest collisions
func (f *Framework) GetClosestCollisonsFor(collision *Collision) []*Collision {
	res := make([]*Collision, 0)
	for _, c := range f.collisions {
		if collision != c {
			res = append(res, c)
		}
	}
	return res
}

func (f *Framework) RegisterCollision(collision *Collision, owner ICollisionOwner) {
	if owner == nil {
		return
	}
	collision.SetEntity(owner)
	f.collisions = append(f.collisions, collision)
}

func (f *Framework) FlushCollisions() {
	f.collisions = make([]*Collision, 0)
}

func (f *Framework) AddAfterUpdate(afterUpdate func()) {
	f.afterUpdates = append(f.afterUpdates, afterUpdate)
}
