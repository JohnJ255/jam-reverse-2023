package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/framework"
)

type Wall struct {
	*framework.GameEntity
	*framework.Sprite
	Position framework.Vec2
	Size     framework.Size
	Rotation framework.Radian
}

func NewWall(pos framework.Vec2, size framework.Size) *Wall {
	w := &Wall{
		Sprite:   framework.InitSprites(size),
		Position: pos,
		Size:     size,
	}
	w.Sprite.Visible = false
	w.GameEntity = framework.InitGameEntity(w)

	return w
}

func (w *Wall) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	op := w.PivotTransform(scale, framework.VecUV{})
	op.GeoM.Rotate(float64(w.Rotation))
	op.GeoM.Translate(w.Position.X, w.Position.Y)

	return op
}

func (w *Wall) GetName() string {
	return "Wall"
}

func (w *Wall) GetModel() framework.Model {
	return nil
}

func (w *Wall) GetSize() framework.Size {
	return w.Size
}

func (w *Wall) GetPosition() framework.Vec2 {
	return w.Position
}

func (w *Wall) GetRotation() framework.Radian {
	return w.Rotation
}

func (w *Wall) SetPosition(pos framework.Vec2) {
	w.Position = pos
}

func (w *Wall) SetRotation(rot framework.Radian) {
	w.Rotation = rot
}

func (w *Wall) IsFixed() bool {
	return true
}

func (w *Wall) GetMass() float64 {
	return 1000
}

func (w *Wall) GetFriction() float64 {
	return 1000
}

func (w *Wall) GetScale() framework.Vec2 {
	return framework.Vec2{1, 1}
}

func (w *Wall) GetPivot() framework.VecUV {
	return framework.VecUV{}
}
