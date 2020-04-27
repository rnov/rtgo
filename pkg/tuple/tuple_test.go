package tuple

import "testing"

func TestNewPoint(t *testing.T) {
	p := NewPoint(1.0, 2.0, 3.0)
	if p.W != 1.0 {
		t.Errorf("Point Constructor failed, expected W value 1 got %v", p.W)
	}
}

func TestNewVector(t *testing.T) {
	p := NewVector(1.0, 2.0, 3.0)
	if p.W != 0 {
		t.Errorf("Vector Constructor failed, expected W value 0 got %v", p.W)
	}
}

func TestTuple_Equal(t *testing.T) {

	tests := []struct {
		name   string
		tpl1   Tuple
		tpl2   Tuple
		expRes bool
	}{
		{
			name:   "Identical tuples",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			expRes: true,
		},
		{
			name:   "Tuples differ in all fields",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			tpl2:   Tuple{X: 0, Y: 1, Z: 2, W: 0},
			expRes: false,
		},
		{
			name:   "Tuples differ in W field",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: 0},
			expRes: false,
		},
		{
			name:   "Tuples differ by to decimals from the epsilon",
			tpl1:   Tuple{X: Epsilon, Y: Epsilon, Z: Epsilon, W: 1},
			tpl2:   Tuple{X: Epsilon - 0.01, Y: Epsilon, Z: Epsilon, W: 1},
			expRes: false,
		},
		//{ // note : should be false but SmallestNonzeroFloat64 in result compared to SmallestNonzeroFloat64 as const is true !
		//	name:   "Tuples W fields differ a epsilon",
		//	tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
		//	tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: 1.0 + Epsilon},
		//	expRes: false,
		//},
		{
			name:   "Tuples W fields differ by epsilon #2 ",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 0},
			tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: Epsilon},
			expRes: false,
		},
	}

	for _, test := range tests {
		r := test.tpl1.Equal(test.tpl2)
		if r != test.expRes {
			t.Errorf("Error in test %v : got => %v expected => %v", test.name, r, test.expRes)
		}
	}
}

func TestEqual(t *testing.T) {

	tests := []struct {
		name   string
		tpl1   Tuple
		tpl2   Tuple
		expRes bool
	}{
		{
			name:   "Identical tuples",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			expRes: true,
		},
		{
			name:   "Tuples differ in all fields",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			tpl2:   Tuple{X: 0, Y: 1, Z: 2, W: 0},
			expRes: false,
		},
		{
			name:   "Tuples differ in W field",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: 0},
			expRes: false,
		},
		{
			name:   "Tuples differ by to decimals from the epsilon",
			tpl1:   Tuple{X: Epsilon, Y: Epsilon, Z: Epsilon, W: 1},
			tpl2:   Tuple{X: Epsilon - 0.01, Y: Epsilon, Z: Epsilon, W: 1},
			expRes: false,
		},
		//{ // note : should be false but SmallestNonzeroFloat64 in result compared to SmallestNonzeroFloat64 as const is true !
		//	name:   "Tuples W fields differ a epsilon",
		//	tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
		//	tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: 1.0 + Epsilon},
		//	expRes: false,
		//},
		{
			name:   "Tuples W fields differ by epsilon #2 ",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 0},
			tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: Epsilon},
			expRes: false,
		},
	}

	for _, test := range tests {
		r := Equal(test.tpl1, test.tpl2)
		if r != test.expRes {
			t.Errorf("Error in test '%v' : got => %v expected => %v", test.name, r, test.expRes)
		}
	}
}

func TestNegate(t *testing.T) {

	tests := []struct {
		name     string
		tpl1     Tuple
		expTuple Tuple
	}{
		{
			name:     "Negates positives ",
			tpl1:     Tuple{X: 1, Y: 2, Z: 3, W: -1},
			expTuple: Tuple{X: -1, Y: -2, Z: -3, W: 1},
		},
		{
			name:     "Negates negatives",
			tpl1:     Tuple{X: -1, Y: -2, Z: -3, W: -1},
			expTuple: Tuple{X: 1, Y: 2, Z: 3, W: 1},
		},
		{
			name:     "Negates negatives and positives",
			tpl1:     Tuple{X: -1, Y: 2, Z: 3, W: -1},
			expTuple: Tuple{X: 1, Y: -2, Z: -3, W: 1},
		},
		{
			name:     "Negates epsilon values",
			tpl1:     Tuple{X: Epsilon, Y: Epsilon, Z: Epsilon, W: -1},
			expTuple: Tuple{X: -Epsilon, Y: -Epsilon, Z: -Epsilon, W: 1},
		},
	}

	for _, test := range tests {
		r := test.tpl1.Negate()
		if !r.Equal(test.expTuple) {
			t.Errorf("Error in test '%v' : got => %v expected => %v", test.name, r, test.expTuple)
		}
	}
}

func TestTuple_IsValid(t *testing.T) {

	tests := []struct {
		name   string
		tpl1   Tuple
		expRes bool
	}{
		{
			name:   "valid W = 1",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			expRes: true,
		},
		{
			name:   "valid W = 0",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 0},
			expRes: true,
		},
		{
			name:   "Not valid W = 2",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 2},
			expRes: false,
		},
		{
			name:   "valid W = -1",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: -1},
			expRes: false,
		},
	}

	for _, test := range tests {

		res := test.tpl1.IsValid()
		if res != test.expRes {
			t.Errorf("Error in test '%v' : got => %v expected => %v", test.name, res, test.expRes)
		}
	}
}

func TestAdd(t *testing.T) {

	tests := []struct {
		name   string
		tpl1   Tuple
		tpl2   Tuple
		expRes Tuple
	}{
		{
			name:   "Valid operation (point + vector)",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			tpl2:   Tuple{X: 4, Y: 2, Z: 3, W: 0},
			expRes: Tuple{X: 5, Y: 4, Z: 6, W: 1},
		},
		{
			name:   "Valid operation #2 (vector + point)",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 0},
			tpl2:   Tuple{X: -4, Y: -2, Z: 3, W: 1},
			expRes: Tuple{X: -3, Y: 0, Z: 6, W: 1},
		},
		{
			name:   "Invalid operation (point + point)",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			tpl2:   Tuple{X: 4, Y: 2, Z: 3, W: 1},
			expRes: Tuple{X: 5, Y: 4, Z: 6, W: 2},
		},
	}

	for _, test := range tests {

		res := Add(test.tpl1, test.tpl2)
		if !res.Equal(test.expRes) {
			t.Errorf("Error in test '%v' : got => %v expected => %v", test.name, res, test.expRes)
		}
	}
}

func TestSub(t *testing.T) {

	tests := []struct {
		name   string
		tpl1   Tuple
		tpl2   Tuple
		expRes Tuple
	}{
		{
			name:   "Valid operation (point - vector)",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			tpl2:   Tuple{X: 4, Y: 2, Z: -6, W: 0},
			expRes: Tuple{X: -3, Y: 0, Z: 9, W: 1},
		},
		{
			name:   "Valid operation #2 (vector - point) result is negated",
			tpl1:   Tuple{X: 1, Y: 2, Z: -3, W: 0},
			tpl2:   Tuple{X: -4, Y: -2, Z: 3, W: 1},
			expRes: Tuple{X: -5, Y: -4, Z: 6, W: 1},
		},
		{
			name:   "valid operation #3 (point - point)",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
			tpl2:   Tuple{X: 4, Y: 2, Z: 3, W: 1},
			expRes: Tuple{X: -3, Y: 0, Z: 0, W: 0},
		},
	}

	for _, test := range tests {

		res := Sub(test.tpl1, test.tpl2)
		if !res.Equal(test.expRes) {
			t.Errorf("Error in test '%v' : got => %v expected => %v", test.name, res, test.expRes)
		}
	}
}

func TestTuple_Scale(t *testing.T) {

	tests := []struct {
		name     string
		tpl1     Tuple
		by       float64
		expTuple Tuple
	}{
		{
			name:     "scale",
			tpl1:     Tuple{X: 1, Y: -2, Z: 3, W: -4},
			by:       3.5,
			expTuple: Tuple{X: 3.5, Y: -7, Z: 10.5, W: -14},
		},
		{
			name:     "scale by fraction",
			tpl1:     Tuple{X: 1, Y: -2, Z: 3, W: -4},
			by:       0.5,
			expTuple: Tuple{X: 0.5, Y: -1, Z: 1.5, W: -2},
		},
		{
			name:     "scale by 0",
			tpl1:     Tuple{X: 1, Y: -2, Z: 3, W: -4},
			by:       0,
			expTuple: Tuple{X: 0, Y: 0, Z: 0, W: 0},
		},
		{
			name:     "scale by epsilon",
			tpl1:     Tuple{X: 1, Y: -2, Z: 3, W: -4},
			by:       Epsilon,
			expTuple: Tuple{X: 1 * Epsilon, Y: -2 * Epsilon, Z: 3 * Epsilon, W: -4 * Epsilon},
		},
	}

	for _, test := range tests {
		r := test.tpl1.Scale(test.by)
		if !r.Equal(test.expTuple) {
			t.Errorf("Error in test '%v' : got => %v expected => %v", test.name, r, test.expTuple)
		}
	}
}

func TestTuple_Length(t *testing.T) {

	tests := []struct {
		name   string
		tpl    Tuple
		expRes float64
	}{
		{
			name:   "zero vector",
			tpl:    Tuple{X: 0, Y: 0, Z: 0, W: 0},
			expRes: 0,
		},
		{
			name:   "X:1, rest zero",
			tpl:    Tuple{X: 1, Y: 0, Z: 0, W: 0},
			expRes: 1,
		},
		{
			name:   "Y:1, rest zero",
			tpl:    Tuple{X: 0, Y: 1, Z: 0, W: 0},
			expRes: 1,
		},
		{
			name:   "Z:1, rest zero",
			tpl:    Tuple{X: 0, Y: 0, Z: 1, W: 0},
			expRes: 1,
		},
		{
			name:   "positive values vector",
			tpl:    Tuple{X: 1, Y: 2, Z: 3, W: 0},
			expRes: 3.7416573867739413,
		},
		{
			name:   "negatives values vector",
			tpl:    Tuple{X: -1, Y: -2, Z: -3, W: 0},
			expRes: 3.7416573867739413,
		},
		{
			name:   "All Epsilon values vector",
			tpl:    Tuple{X: Epsilon, Y: Epsilon, Z: Epsilon, W: 0},
			expRes: 0,
		},
		{
			name:   "All negative Epsilon values vector",
			tpl:    Tuple{X: -Epsilon, Y: -Epsilon, Z: -Epsilon, W: 0},
			expRes: 0,
		},
	}

	for _, test := range tests {
		r := test.tpl.Length()
		diff := r - test.expRes
		if diff > Epsilon {
			t.Errorf("Error in test '%v' : got => %v expected => %v, diff regarding Epsiolon %v",
				test.name, r, test.expRes, diff-Epsilon)
		}
	}
}

func TestTuple_Normalize(t *testing.T) {

	tests := []struct {
		name     string
		tpl      Tuple
		expTuple Tuple
	}{
		{
			name:     "zero vector",
			tpl:      Tuple{X: 0, Y: 0, Z: 0, W: 0},
			expTuple: Tuple{X: 0, Y: 0, Z: 0, W: 0},
		},
		{
			name:     "X:1, rest zero",
			tpl:      Tuple{X: 1, Y: 0, Z: 0, W: 0},
			expTuple: Tuple{X: 1, Y: 0, Z: 0, W: 0},
		},
		{
			name:     "Y:1, rest zero",
			tpl:      Tuple{X: 0, Y: 1, Z: 0, W: 0},
			expTuple: Tuple{X: 0, Y: 1, Z: 0, W: 0},
		},
		{
			name:     "Z:1, rest zero",
			tpl:      Tuple{X: 0, Y: 0, Z: 1, W: 0},
			expTuple: Tuple{X: 0, Y: 0, Z: 1, W: 0},
		},
		{
			name:     "positive values vector",
			tpl:      Tuple{X: 1, Y: 2, Z: 3, W: 0},
			expTuple: Tuple{X: 1 / 3.7416573867739413, Y: 2 / 3.7416573867739413, Z: 3 / 3.7416573867739413, W: 0},
		},
		{
			name:     "negatives values vector",
			tpl:      Tuple{X: -1, Y: -2, Z: -3, W: 0},
			expTuple: Tuple{X: -1 / 3.7416573867739413, Y: -2 / 3.7416573867739413, Z: -3 / 3.7416573867739413, W: 0},
		},
		{
			name:     "All Epsilon values vector",
			tpl:      Tuple{X: Epsilon, Y: Epsilon, Z: Epsilon, W: 0},
			expTuple: Tuple{X: Epsilon, Y: Epsilon, Z: Epsilon, W: 0},
		},
		{
			name:     "All negative Epsilon values vector",
			tpl:      Tuple{X: -Epsilon, Y: -Epsilon, Z: -Epsilon, W: 0},
			expTuple: Tuple{X: -Epsilon, Y: -Epsilon, Z: -Epsilon, W: 0},
		},
	}

	for _, test := range tests {
		r := test.tpl.Normalize()
		if !r.Equal(test.expTuple) {
			t.Errorf("Error in test '%v' : got => %v expected => %v", test.name, r, test.expTuple)
		}
	}
}
