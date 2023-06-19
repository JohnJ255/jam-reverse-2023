package framework

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"strings"
	"time"
)

type Game interface {
	Start(f *Framework)
	Update(dt float64) error
	Draw(screen *ebiten.Image)
}

type Framework struct {
	game         Game
	lastUpdate   time.Time
	console      *Console
	windowWidth  int
	windowHeight int
	windowTitle  string
	ticks        uint64
	debugDraws   map[string]func(screen *ebiten.Image)
}

var fw *Framework

func InitWindowGame(g Game, windowWidth, windowHeight int, windowTitle string) *Framework {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle(windowTitle)
	fw = &Framework{
		game:         g,
		windowWidth:  windowWidth,
		windowHeight: windowHeight,
		windowTitle:  windowTitle,
		console:      NewConsole(),
		debugDraws:   make(map[string]func(screen *ebiten.Image)),
	}
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
	if f.ticks == 1 {
		f.game.Start(f)
		return nil
	}

	if f.console.IsAvailable && inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		f.console.Toggle()
	}

	if f.console.IsOpened {
		f.console.Update(f)

		return nil
	}

	return f.game.Update(dt)
}

func (f *Framework) Draw(screen *ebiten.Image) {
	f.game.Draw(screen)
	if f.console.IsOpened {
		f.console.Draw(screen, 0, 0, f.windowWidth, f.windowHeight/3)
	}
	for _, drawer := range f.debugDraws {
		drawer(screen)
	}
}

func (g *Framework) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
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

func (f *Framework) SetDebugDraw(name string, drawer func(screen *ebiten.Image)) {
	f.debugDraws[name] = drawer
}

func (f *Framework) RemoveDebugDraw(name string) {
	delete(f.debugDraws, name)
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
