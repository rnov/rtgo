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

func Equal(a, b Tuple) bool {
	return math.Abs(a.X-b.X) < Epsilon &&
		math.Abs(a.Y-b.Y) < Epsilon &&
		math.Abs(a.Z-b.Z) < Epsilon &&
		math.Abs(a.W-b.W) < Epsilon
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

func Add(a, b Tuple) Tuple {
	return Tuple{
		X: a.X + b.X,
		Y: a.Y + b.Y,
		Z: a.Z + b.Z,
		W: a.W + b.W,
	}
}

func Sub(a, b Tuple) Tuple {

	t := Tuple{
		X: a.X - b.X,
		Y: a.Y - b.Y,
		Z: a.Z - b.Z,
		W: a.W - b.W,
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

// It does make sense to apply it only to vector, W is included in case is misused
func DotProduct(a, b Tuple) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

// It does make sense to apply it only to vector
func CrossProduct(a, b Tuple) Tuple {
	return Tuple{
		X: a.Y*b.Z - a.Z*b.Y,
		Y: a.Z*b.X - a.X*b.Z,
		Z: a.X*b.Y - a.Y*b.X,
	}
}
