package game

import "image"

func area(a, b, c image.Point) int {
	return (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
}

func intersect1(a, b, c, d int) bool {
	if a > b {
		a, b = b, a
	}
	if c > d {
		c, d = d, c
	}

	maxac := a
	if a < c {
		maxac = c
	}
	minbd := b
	if b > d {
		minbd = d
	}
	return maxac <= minbd
}

func intersect(a, b, c, d image.Point) bool {
	return intersect1(a.X, b.X, c.X, d.X) &&
		intersect1(a.Y, b.Y, c.Y, d.Y) &&
		area(a, b, c)*area(a, b, d) < 0 &&
		area(c, d, a)*area(c, d, b) < 0
}
