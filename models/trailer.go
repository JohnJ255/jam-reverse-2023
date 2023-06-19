package models

import (
	"math"
	"reverse-jam-2023/helper"
)

type TowbarInterface interface {
	GetPosition() helper.DirectionPosition
	GetTowbarPosition() helper.Position
}

type TrailerType int

const (
	TrailerTypeNone TrailerType = iota
	TrailerTypeCart
	TrailerTypeTrailer
)

type Trailer struct {
	Trailer      TrailerJoin
	Traktor      TowbarInterface
	Position     helper.DirectionPosition
	Size         helper.Size
	Pivot        helper.PositionUV
	trType       TrailerType
	health       int
	maxHealth    int
	baseInertion float64
	mass         float64
}

func NewTrailer(size helper.Size, mass float64, trType TrailerType) *Trailer {
	return &Trailer{
		Size:         size,
		Pivot:        helper.PositionUV{1, 0.5},
		maxHealth:    100,
		health:       100,
		mass:         mass,
		trType:       trType,
		baseInertion: 0.94,
	}
}

func NewTrailerToBackOfTractor(trac TowbarInterface, size helper.Size, mass float64, trType TrailerType) *Trailer {
	pos := trac.GetPosition()
	pos.X = trac.GetTowbarPosition().X
	pos.Y = trac.GetTowbarPosition().Y
	t := NewTrailer(size, mass, trType)
	t.Position = pos
	return t
}

func (t *Trailer) getSelfMass() float64 {
	return t.mass
}

func (t *Trailer) getFullMass() float64 {
	if t.Trailer != nil {
		return t.mass + t.Trailer.getFullMass()
	}

	return t.getSelfMass()
}

func (t *Trailer) getFrictionForce() float64 {
	return 1 - t.calcInertionDependsMass()
}

func (t *Trailer) GetPivot() helper.PositionUV {
	return t.Pivot
}

func (t *Trailer) calcInertionDependsMass() float64 {
	k := 1 + (massEtalon-t.mass)/massEtalon
	return helper.Limited(t.baseInertion-k/10, 0.9, 0.999)
}

func (t *Trailer) AddTraktor(c TowbarInterface) {
	t.Traktor = c
}

func (t *Trailer) GetTowbarLocalPosition() helper.Position {
	x := t.Size.Length * t.Pivot.U * math.Cos(t.Position.Angle)
	y := t.Size.Width * t.Pivot.V * math.Sin(t.Position.Angle)
	return helper.Position{x, y}
}
