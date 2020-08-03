package game

import (
	"image"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIntersect(t *testing.T) {
	p1 := image.Point{X: 100, Y: 100}
	p2 := image.Point{X: 200, Y: 200}
	p3 := image.Point{X: 100, Y: 200}
	p4 := image.Point{X: 200, Y: 100}

	require.True(t, intersect(p1, p2, p3, p4))
}

func TestIntersectNegative(t *testing.T) {
	p1 := image.Point{X: 100, Y: 100}
	p2 := image.Point{X: 100, Y: 200}
	p3 := image.Point{X: 200, Y: 200}
	p4 := image.Point{X: 200, Y: 100}

	require.False(t, intersect(p1, p2, p3, p4))
}

func TestIntersectSame(t *testing.T) {
	p1 := image.Point{X: 100, Y: 100}
	p2 := image.Point{X: 100, Y: 200}
	p3 := image.Point{X: 200, Y: 200}

	require.False(t, intersect(p1, p2, p3, p2))
}
