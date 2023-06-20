package helper

import "math"

type Size struct {
	Width, Length float64
}

type IntSize struct {
	Width, Height int
}

type DirectionPosition struct {
	Vec2
	Angle Radian
}

type Vec2 struct {
	X, Y float64
}

type VecUV struct {
	U, V float64
}

type Degrees float64
type Radian float64

func (d Degrees) ToRadians() Radian {
	return Radian(float64(d) * math.Pi / 180)
}

func (r Radian) ToVec() Vec2 {
	return Vec2{math.Cos(float64(r)), math.Sin(float64(r))}
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

func (v Vec2) Mul(n float64) Vec2 {
	return Vec2{v.X * n, v.Y * n}
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Normalize() Vec2 {
	if v.X == 0 && v.Y == 0 {
		return v
	}
	return v.Mul(1 / v.Length())
}

func (v Vec2) ToRadian() Radian {
	return Radian(math.Atan2(v.Y, v.X))
}
