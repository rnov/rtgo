package tuple

import (
	"math"
)

const (
	EPSILON = math.SmallestNonzeroFloat64
)

type Tuple struct {
	X, Y, Z, W float64
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 1.0,
	}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 0,
	}
}

func (t Tuple) Compare(to Tuple) bool {
	return math.Abs(t.X-to.X) < EPSILON &&
		math.Abs(t.Y-to.Y) < EPSILON &&
		math.Abs(t.Z-to.Z) < EPSILON &&
		math.Abs(t.W-to.W) < EPSILON
}

func Compare(t, t2 Tuple) bool {
	return math.Abs(t.X-t2.X) < EPSILON &&
		math.Abs(t.Y-t2.Y) < EPSILON &&
		math.Abs(t.Z-t2.Z) < EPSILON &&
		math.Abs(t.W-t2.W) < EPSILON
}
