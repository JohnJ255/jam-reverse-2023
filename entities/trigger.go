package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/framework"
)

type Trigger struct {
	*framework.GameEntity
	*framework.Sprite
	Position      framework.Vec2
	Size          framework.Size
	Rotation      framework.Radian
	TriggerAction func(entity framework.ICollisionOwner, collide *framework.Collide)
}

func NewTrigger(pos framework.Vec2, size framework.Size, onTrigger func(entity framework.ICollisionOwner, collide *framework.Collide)) *Trigger {
	t := &Trigger{
		Sprite:        framework.InitSprites(size),
		Position:      pos,
		Size:          size,
		TriggerAction: onTrigger,
	}
	t.Sprite.Visible = false

	t.GameEntity = framework.InitGameEntity(t)

	return t
}

func (t *Trigger) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	op := t.PivotTransform(scale, framework.VecUV{})
	op.GeoM.Rotate(float64(t.Rotation))
	op.GeoM.Translate(t.Position.X, t.Position.Y)

	return op
}

func (t *Trigger) OnTrigger(entity framework.ICollisionOwner, collide *framework.Collide) {
	t.TriggerAction(entity, collide)
}

func (t *Trigger) GetName() string {
	return "Trigger"
}

func (t *Trigger) GetModel() framework.Model {
	return nil
}

func (t *Trigger) GetSize() framework.Size {
	return t.Size
}

func (t *Trigger) GetPosition() framework.Vec2 {
	return t.Position
}

func (t *Trigger) GetRotation() framework.Radian {
	return t.Rotation
}

func (t *Trigger) SetPosition(pos framework.Vec2) {
	t.Position = pos
}

func (t *Trigger) SetRotation(rot framework.Radian) {
	t.Rotation = rot
}

func (t *Trigger) GetScale() framework.Vec2 {
	return framework.Vec2{1, 1}
}

func (t *Trigger) GetPivot() framework.VecUV {
	return framework.VecUV{}
}
