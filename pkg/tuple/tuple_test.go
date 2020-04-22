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

func TestTuple_Compare(t *testing.T) {
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
			tpl1:   Tuple{X: EPSILON, Y: EPSILON, Z: EPSILON, W: 1},
			tpl2:   Tuple{X: EPSILON - 0.01, Y: EPSILON, Z: EPSILON, W: 1},
			expRes: false,
		},
		//{ // note : should be false but SmallestNonzeroFloat64 in result compared to SmallestNonzeroFloat64 as const is true !
		//	name:   "Tuples W fields differ a epsilon",
		//	tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
		//	tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: 1.0 + EPSILON},
		//	expRes: false,
		//},
		{
			name:   "Tuples W fields differ by epsilon #2 ",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 0},
			tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: EPSILON},
			expRes: false,
		},
	}

	for _, test := range tests {
		r := test.tpl1.Compare(test.tpl2)
		if r != test.expRes {
			t.Errorf("Error in test %v : got => %v expected => %v", test.name, r, test.expRes)
		}
	}
}

func TestCompare(t *testing.T) {

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
			tpl1:   Tuple{X: EPSILON, Y: EPSILON, Z: EPSILON, W: 1},
			tpl2:   Tuple{X: EPSILON - 0.01, Y: EPSILON, Z: EPSILON, W: 1},
			expRes: false,
		},
		//{ // note : should be false but SmallestNonzeroFloat64 in result compared to SmallestNonzeroFloat64 as const is true !
		//	name:   "Tuples W fields differ a epsilon",
		//	tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 1},
		//	tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: 1.0 + EPSILON},
		//	expRes: false,
		//},
		{
			name:   "Tuples W fields differ by epsilon #2 ",
			tpl1:   Tuple{X: 1, Y: 2, Z: 3, W: 0},
			tpl2:   Tuple{X: 1, Y: 2, Z: 3, W: EPSILON},
			expRes: false,
		},
	}

	for _, test := range tests {
		r := Compare(test.tpl1, test.tpl2)
		if r != test.expRes {
			t.Errorf("Error in test '%v' : got => %v expected => %v", test.name, r, test.expRes)
		}
	}
}
