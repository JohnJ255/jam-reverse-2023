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

func NewTrailer(pos framework.DirectionPosition, size framework.Size, mass float64, trType models.TrailerType) *TrailerEntity {
	t := &TrailerEntity{
		Sprite:  framework.InitSprites(size),
		Trailer: models.NewTrailer(size, mass, trType),
	}
	t.GameEntity = framework.InitGameEntity(t)
	t.Trailer.Position = pos
	t.LoadResources(&loader.ResourceLoader{}, loader.TrailerFileNames[trType])

	return t
}

func NewTrailerToBackOfTractor(trac models.TowbarInterface, size framework.Size, mass float64, trType models.TrailerType) *TrailerEntity {
	pos := trac.GetPosition()
	t := NewTrailer(pos, size, mass, trType)
	t.Trailer.Position.X = trac.GetTowbarPosition().X - t.Trailer.GetTowbarLocalPosition().X
	t.Trailer.Position.Y = trac.GetTowbarPosition().Y - t.Trailer.GetTowbarLocalPosition().Y
	return t
}

func (t *TrailerEntity) GetTransforms(scale float64) *ebiten.DrawImageOptions {
	op := t.PivotTransform(scale, t.Trailer.Pivot)
	op.GeoM.Rotate(float64(t.Trailer.Position.Angle))
	op.GeoM.Translate(t.Trailer.Position.X, t.Trailer.Position.Y)

	return op
}

func (t *TrailerEntity) GetSize() framework.Size {
	return t.Trailer.Size
}

func (t *TrailerEntity) GetPosition() framework.Vec2 {
	return t.Trailer.Position.Vec2
}

func (t *TrailerEntity) SetPosition(pos framework.Vec2) {
	t.Trailer.Position.Vec2 = pos
}

func (t *TrailerEntity) GetRotation() framework.Radian {
	return t.Trailer.Position.Angle
}

func (t *TrailerEntity) SetRotation(r framework.Radian) {
	t.Trailer.Position.Angle = r
}

func (t *TrailerEntity) GetScale() framework.Vec2 {
	return framework.Vec2{1, 1}
}

func (t *TrailerEntity) GetPivot() framework.VecUV {
	return t.Trailer.Pivot
}

func (t *TrailerEntity) GetName() string {
	return "trailer"
}

func (t *TrailerEntity) GetModel() framework.Model {
	return t.Trailer
}

func (t *TrailerEntity) Update(dt float64) {
	t.GameEntity.Update(dt)
	t.Trailer.Control()
}

func (t *TrailerEntity) IsFixed() bool {
	return false
}

func (t *TrailerEntity) GetMass() float64 {
	return t.Trailer.GetMass()
}

func (t *TrailerEntity) GetFriction() float64 {
	return 1
}
