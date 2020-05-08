package canvas

type color struct {
	r float64
	g float64
	b float64
}

func NewColor(r, g, b float64) color {
	return color{
		r: r,
		g: g,
		b: b,
	}
}

func add(c1, c2 color) color {
	// todo
	return color{}
}

func sub(c1, c2 color) color {
	// todo
	return color{}
}

func (c *color) multByScalar() color {
	// todo
	return color{}
}

// A tuple multiplication
func hadamardProdct(c1, c2 color) color {
	// todo
	return color{}
}
