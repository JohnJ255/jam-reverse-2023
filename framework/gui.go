package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font"
	"image/color"
)

type Button struct {
	IsVisible   bool
	Text        string
	OnClick     func()
	FillColor   color.Color
	StrokeColor color.Color
	position    Vec2
	size        Size
	fontFace    font.Face
}

func NewButton(text string, onClick func(), fontFace font.Face) *Button {
	return &Button{
		IsVisible:   true,
		Text:        text,
		OnClick:     onClick,
		FillColor:   color.NRGBA{100, 100, 100, 255},
		StrokeColor: color.White,
		fontFace:    fontFace,
	}
}

func (b *Button) Draw(screen *ebiten.Image, x, y, w, h int) {
	b.position = Vec2{float64(x), float64(y)}
	b.size = Size{float64(w), float64(h)}
	vector.DrawFilledRect(screen, float32(x), float32(y), float32(w), float32(h), b.FillColor, false)
	vector.StrokeRect(screen, float32(x), float32(y), float32(w), float32(h), 2, b.StrokeColor, false)
	text.Draw(screen, b.Text, b.fontFace, x+w/2-len(b.Text)*3, y+h/2+5, color.NRGBA{230, 240, 250, 255})
}

func (b *Button) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		pos := Vec2{float64(x), float64(y)}
		if pos.InRectWithoutRotation(b.position, b.size) {
			b.OnClick()
		}
	}
}
