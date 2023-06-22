package framework

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ICollisionFiguresContainer interface {
	GetFigures() []ICollisionFigure
}

type ICollisionPainter interface {
	Draw(screen *ebiten.Image, figure ICollisionFigure)
}

type IIntersectionPainter interface {
	Draw(screen *ebiten.Image, intersection ContactSet)
}

type DebugTool struct {
	f                  *Framework
	Draws              map[string]func(screen *ebiten.Image)
	CollisionDrawer    ICollisionPainter
	IntersectionDrawer IIntersectionPainter
}

func NewDebugTool(f *Framework, drawer ICollisionPainter, interDrawer IIntersectionPainter) *DebugTool {
	return &DebugTool{
		f:                  f,
		Draws:              make(map[string]func(screen *ebiten.Image)),
		CollisionDrawer:    drawer,
		IntersectionDrawer: interDrawer,
	}
}

func (t *DebugTool) DefaultDrawCollisions(screen *ebiten.Image) {
	for _, ent := range t.f.entities {
		for _, cmp := range ent.GetComponents() {
			switch cmp.(type) {
			case ICollisionFiguresContainer:
				for _, fig := range cmp.(ICollisionFiguresContainer).GetFigures() {
					t.CollisionDrawer.Draw(screen, fig)
				}
			}
		}
	}
}

func (t *DebugTool) DefaultDrawIntersections(csList []ContactSet) func(screen *ebiten.Image) {
	return func(screen *ebiten.Image) {
		for _, cs := range csList {
			t.IntersectionDrawer.Draw(screen, cs)
		}
	}
}

func (t *DebugTool) SetDebugDraw(name string, drawer func(screen *ebiten.Image)) {
	t.Draws[name] = drawer
}

func (t *DebugTool) RemoveDebugDraw(name string) {
	delete(t.Draws, name)
}
