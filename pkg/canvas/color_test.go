package canvas

import (
	"testing"

	tuple "github.com/rnov/rtgo/pkg/tuple"
)

func TestNewColor(t *testing.T) {

}

func TestAdd(t *testing.T) {

	tests := []struct {
		name   string
		tplA   color
		tplB   color
		expRes color
	}{
		{
			name:   "c1 + c2 operation",
			tplA:   color{r: 0.1, g: 0.2, b: 0.3},
			tplB:   color{r: 0.4, g: 0.2, b: 0.3},
			expRes: color{r: 0.3, g: 0.4, b: 0.6},
		},
		{
			name:   "zero value operation",
			tplA:   color{r: 0.1, g: 0.2, b: 0.3},
			tplB:   color{r: 0, g: 0, b: 0},
			expRes: color{r: 0.1, g: 0.2, b: 0.3},
		},
		{
			name:   "Epsilon value operation",
			tplA:   color{r: 0.1, g: 0.2, b: 0.3},
			tplB:   color{r: tuple.Epsilon, g: tuple.Epsilon, b: tuple.Epsilon},
			expRes: color{r: 0.1 + tuple.Epsilon, g: 0.2 + tuple.Epsilon, b: 0.3 + tuple.Epsilon},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//	todo
		})
	}
}

func TestSub(t *testing.T) {

	tests := []struct {
		name   string
		tplA   color
		tplB   color
		expRes color
	}{
		{
			name:   "operation#1 c1 + c2",
			tplA:   color{r: 0.1, g: 0.2, b: 0.3},
			tplB:   color{r: 0.4, g: 0.2, b: 0.3},
			expRes: color{r: 0.3, g: 0.4, b: 0.6},
		},
		{
			name:   "operation#2 c1 + c2",
			tplA:   color{r: 0.1, g: 0.2, b: 0.3},
			tplB:   color{r: 0, g: 0, b: 0},
			expRes: color{r: 0.1, g: 0.2, b: 0.3},
		},
		{
			name:   "operation#3 c1 + c2",
			tplA:   color{r: 0.1, g: 0.2, b: 0.3},
			tplB:   color{r: tuple.Epsilon, g: tuple.Epsilon, b: tuple.Epsilon},
			expRes: color{r: 0.1, g: 0.2, b: 0.3},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//	todo
		})
	}
}

func TestMultByScalar(t *testing.T) {

	tests := []struct {
		name     string
		tplA     color
		by       float64
		expColor color
	}{
		{
			name:     "multiply",
			tplA:     color{r: 0.1, g: 0.2, b: 0.3},
			by:       3.5,
			expColor: color{r: 0.35, g: 0.7, b: 1.05},
		},
		{
			name:     "multiply by 0",
			tplA:     color{r: 0.1, g: 0.2, b: 0.3},
			by:       0,
			expColor: color{r: 0, g: 0, b: 0},
		},
		{
			name:     "multiply by epsilon",
			tplA:     color{r: 0.1, g: 0.2, b: 0.3},
			by:       tuple.Epsilon,
			expColor: color{r: 0.1 * tuple.Epsilon, g: 0.2 * tuple.Epsilon, b: 0.3 * tuple.Epsilon},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//	todo
		})
	}
}

func TestHadamardProdct(t *testing.T) {

	tests := []struct {
		name     string
		tplA     color
		tplB     color
		expColor color
	}{
		{
			name:     "zero value color product",
			tplA:     color{r: 0, g: 0, b: 0},
			tplB:     color{r: 0.4, g: 0.2, b: 0.3},
			expColor: color{r: 0, g: 0, b: 0},
		},
		{
			name:     "epsilon value color product",
			tplA:     color{r: tuple.Epsilon, g: tuple.Epsilon, b: tuple.Epsilon},
			tplB:     color{r: 1, g: 1, b: 1},
			expColor: color{r: 1, g: 1, b: 1},
		},
		{
			name:     "value colors product",
			tplA:     color{r: 0.1, g: 0.2, b: 0.3},
			tplB:     color{r: 0.4, g: 0.2, b: 0.3},
			expColor: color{r: 0.4, g: 0.04, b: 0.09},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			//	todo
		})
	}
}
