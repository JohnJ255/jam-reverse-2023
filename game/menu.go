package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"reverse-jam-2023/framework"
)

type Menu struct {
	IsOpened bool
	game     *Game
	buttons  []framework.IGUIElement
}

func NewMenu(game *Game) *Menu {
	m := &Menu{
		game: game,
	}
	m.buttons = []framework.IGUIElement{
		framework.NewButton("New game", func() {
			m.game.menu.IsOpened = false
			m.game.level.Change(m.game.f, 1)
		}, m.game.fontGUI),
		framework.NewButton("Restart level", func() {
			m.game.menu.IsOpened = false
			m.game.level.Change(m.game.f, m.game.level.index)
		}, m.game.fontGUI),
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
		}, m.game.fontGUI),
		framework.NewHorizontalPanel([]framework.IGUIElement{
			framework.NewButton("-", func() {
				m.game.SoundMasterVolume = framework.Limited(m.game.SoundMasterVolume-0.1, 0, 1)
				m.game.f.Audio.SetMasterVolume(m.game.SoundMasterVolume)
			}, m.game.fontGUI),
			framework.NewLabel("Sound", m.game.fontGUI),
			framework.NewButton("+", func() {
				m.game.SoundMasterVolume = framework.Limited(m.game.SoundMasterVolume+0.1, 0, 1)
				m.game.f.Audio.SetMasterVolume(m.game.SoundMasterVolume)
			}, m.game.fontGUI),
		}, []float64{20, 60, 20}),
		framework.NewButton("About", func() {
			m.game.menu.IsOpened = false
			m.game.level.Change(m.game.f, LastLevelIndex)
		}, m.game.fontGUI),
	}
	return m
}

func (m *Menu) Draw(screen *ebiten.Image) {
	x := 300
	y := 120
	dy := 40
	w := 140
	h := 30

	vector.DrawFilledRect(screen, float32(x-50), float32(y-50), float32(w+100), float32(4*(h+dy)), color.NRGBA{100, 100, 100, 150}, false)
	text.Draw(screen, m.game.Name, m.game.fontGUI, x+w/2-len(m.game.Name)*4, y-25, color.White)

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
