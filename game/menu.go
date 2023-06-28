package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"reverse-jam-2023/framework"
)

type Menu struct {
	IsOpened bool
	game     *Game
	buttons  []*framework.Button
}

func NewMenu(game *Game) *Menu {
	m := &Menu{
		game: game,
	}
	m.buttons = []*framework.Button{
		framework.NewButton("New game", func() {
			m.game.menu.IsOpened = false
			m.game.level.Change(m.game.f, 1)
		}),
		framework.NewButton("Restart level", func() {
			m.game.menu.IsOpened = false
			m.game.level.Change(m.game.f, m.game.level.index)
		}),
		framework.NewButton("Next level", func() {
			m.game.menu.IsOpened = false
			if m.game.level.index == LastLevelIndex || m.game.level.index == 0 {
				m.game.menu.IsOpened = false
				m.game.level.Change(m.game.f, 1)
			} else {
				m.game.level.Change(m.game.f, m.game.level.index+1)
				if m.game.level.index < LastLevelIndex {
					m.game.level.Score -= NewLevelScore
				}
			}
		}),
		framework.NewButton("Sound >>", func() {
			m.game.SoundMasterVolume = framework.Limited(m.game.SoundMasterVolume+0.1, 0, 1)
			m.game.f.Audio.SetMasterVolume(m.game.SoundMasterVolume)
		}),
		framework.NewButton("<< Sound", func() {
			m.game.SoundMasterVolume = framework.Limited(m.game.SoundMasterVolume-0.1, 0, 1)
			m.game.f.Audio.SetMasterVolume(m.game.SoundMasterVolume)
		}),
		framework.NewButton("About", func() {
			m.game.menu.IsOpened = false
			m.game.level.Change(m.game.f, LastLevelIndex)
		}),
	}
	return m
}

func (m *Menu) Draw(screen *ebiten.Image) {
	x := 300
	y := 100
	dy := 40
	w := 100
	h := 30

	vector.DrawFilledRect(screen, float32(x-30), float32(y-30), float32(w+60), float32(4*(h+dy)), color.NRGBA{100, 100, 100, 150}, false)
	ebitenutil.DebugPrintAt(screen, m.game.Name, x+w/2-len(m.game.Name)*3, y-30)

	for _, b := range m.buttons {
		b.Draw(screen, x, y, w, h)
		y += dy
	}
}

func (m *Menu) Update() {
	for _, b := range m.buttons {
		b.Update()
	}
}
