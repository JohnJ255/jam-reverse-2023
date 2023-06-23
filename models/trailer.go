package models

import (
	"math"
	"reverse-jam-2023/framework"
)

type TowbarInterface interface {
	GetPosition() framework.DirectionPosition
	GetTowbarPosition() framework.Vec2
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
	Position     framework.DirectionPosition
	Size         framework.Size
	Pivot        framework.VecUV
	trType       TrailerType
	health       int
	maxHealth    int
	baseInertion float64
	mass         float64
}

func NewTrailer(size framework.Size, mass float64, trType TrailerType) *Trailer {
	return &Trailer{
		Size:         size,
		Pivot:        framework.VecUV{0.2, 0.5},
		maxHealth:    100,
		health:       100,
		mass:         mass,
		trType:       trType,
		baseInertion: 0.94,
	}
}

func NewTrailerToBackOfTractor(trac TowbarInterface, size framework.Size, mass float64, trType TrailerType) *Trailer {
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

func (t *Trailer) GetPivot() framework.VecUV {
	return t.Pivot
}

func (t *Trailer) calcInertionDependsMass() float64 {
	k := 1 + (massEtalon-t.mass)/massEtalon
	return framework.Limited(t.baseInertion-k/10, 0.9, 0.999)
}

func (t *Trailer) ConnectTraktor(c TowbarInterface) {
	t.Traktor = c
}

func (t *Trailer) DisconnectTraktor() {
	t.Traktor = nil
}

func (t *Trailer) GetTowbarLocalPosition() framework.Vec2 {
	x := t.Size.Length * (1 - t.Pivot.U) * math.Cos(float64(t.Position.Angle))
	y := t.Size.Length * (1 - t.Pivot.U) * math.Sin(float64(t.Position.Angle))
	return framework.Vec2{x, y}
}

func (t *Trailer) Control(velocity framework.Vec2) {
	if t.Traktor == nil {
		return
	}

	tlp := t.GetTowbarLocalPosition()
	t.Position.X = t.Traktor.GetTowbarPosition().X - tlp.X
	t.Position.Y = t.Traktor.GetTowbarPosition().Y - tlp.Y
	t.Position.Angle = tlp.Add(velocity).ToRadian()
}
