package helper

type Size struct {
	Width, Height, Length float64
}

type IntSize struct {
	Width, Height int
}

type DirectionPosition struct {
	Position
	Angle float64
}

type Position struct {
	X, Y float64
}

type PositionUV struct {
	U, V float64
}

type Degrees float64
