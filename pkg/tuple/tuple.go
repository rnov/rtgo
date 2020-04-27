package tuple

import (
	"math"
)

const (
	Epsilon = math.SmallestNonzeroFloat64
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

func (t Tuple) Equal(to Tuple) bool {
	return math.Abs(t.X-to.X) < Epsilon &&
		math.Abs(t.Y-to.Y) < Epsilon &&
		math.Abs(t.Z-to.Z) < Epsilon &&
		math.Abs(t.W-to.W) < Epsilon
}

func Equal(t, t2 Tuple) bool {
	return math.Abs(t.X-t2.X) < Epsilon &&
		math.Abs(t.Y-t2.Y) < Epsilon &&
		math.Abs(t.Z-t2.Z) < Epsilon &&
		math.Abs(t.W-t2.W) < Epsilon
}

func (t *Tuple) Negate() *Tuple {
	t.X = -t.X
	t.Y = -t.Y
	t.Z = -t.Z
	t.W = -t.W
	return t
}

func (t Tuple) IsValid() bool {
	return t.W == 1 || t.W == 0
}

func Add(t1, t2 Tuple) Tuple {
	return Tuple{
		X: t1.X + t2.X,
		Y: t1.Y + t2.Y,
		Z: t1.Z + t2.Z,
		W: t1.W + t2.W,
	}
}

func Sub(t1, t2 Tuple) Tuple {

	t := Tuple{
		X: t1.X - t2.X,
		Y: t1.Y - t2.Y,
		Z: t1.Z - t2.Z,
		W: t1.W - t2.W,
	}
	if !t.IsValid() {
		t.Negate()
	}

	return t
}

func (t *Tuple) Scale(by float64) Tuple {
	t.X *= by
	t.Y *= by
	t.Z *= by
	t.W *= by
	return *t
}

func (t Tuple) Length() float64 {
	l := math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
	return l
}

func (t *Tuple) Normalize() *Tuple {

	if l := t.Length(); l != 0 {
		t.X /= l
		t.Y /= l
		t.Z /= l
	}
	return t
}
