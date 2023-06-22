package entities

import (
	"github.com/hajimehoshi/ebiten/v2"
	"reverse-jam-2023/framework"
	"reverse-jam-2023/loader"
	"reverse-jam-2023/models"
)

type TrailerEntity struct {
	*framework.GameEntity
	*framework.Sprite
	Trailer *models.Trailer
}

func (t *TrailerEntity) GetSize() framework.Size {
	return t.Trailer.Size
}

func (t *TrailerEntity) GetPosition() framework.Vec2 {
	return t.Trailer.Position.Vec2
}

func (t *TrailerEntity) GetRotation() framework.Radian {
	return t.Trailer.Position.Angle
}

func (t *TrailerEntity) GetScale() framework.Vec2 {
	return framework.Vec2{1, 1}
}

func (t *TrailerEntity) GetPivot() framework.VecUV {
	return t.Trailer.Pivot
}

func NewTrailer(pos framework.DirectionPosition, size framework.Size, mass float64, trType models.TrailerType) *TrailerEntity {
	t := &TrailerEntity{
		GameEntity: framework.InitGameEntity(),
		Sprite:     framework.InitSprites(),
		Trailer:    models.NewTrailer(size, mass, trType),
	}
	t.Trailer.Position = pos
	t.LoadResources(&loader.ResourceLoader{}, loader.TrailerFileNames[trType])

	return t
}

func NewTrailerToBackOfTractor(trac models.TowbarInterface, size framework.Size, mass float64, trType models.TrailerType) *TrailerEntity {
	pos := trac.GetPosition()
	pos.X = trac.GetTowbarPosition().X
	pos.Y = trac.GetTowbarPosition().Y
	t := NewTrailer(pos, size, mass, trType)
	return t
}

func (t *TrailerEntity) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	op := t.PivotTransform(scale, t.Trailer.Size, t.Trailer.Pivot)
	op.GeoM.Rotate(float64(t.Trailer.Position.Angle))
	op.GeoM.Translate(t.Trailer.Position.X, t.Trailer.Position.Y)

	return op
}

func (t *TrailerEntity) AddComponent(c framework.IComponent) {
	c.SetOwner(t)
	t.GameEntity.AddComponent(c)
}
