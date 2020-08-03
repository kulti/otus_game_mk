package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type board struct {
	pins       []*image.Point
	ropes      []rope
	touchedPin *image.Point
}

type rope struct {
	a, b *image.Point
	clr  color.Color
}

func newBoard() *board {
	ropes, pins := loadLevel()
	b := &board{
		pins:  pins,
		ropes: ropes,
	}

	b.updateRopes()

	return b
}

func (b *board) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if b.touchedPin != nil {
			b.touchedPin.X = x
			b.touchedPin.Y = y
		} else {
			b.touchedPin = b.findPin(x, y)
		}
	} else {
		b.touchedPin = nil
	}

	b.updateRopes()

	b.dumpLevel()
}

func (b *board) findPin(x, y int) *image.Point {
	for _, p := range b.pins {
		if isNearPin(p, x, y) {
			return p
		}
	}
	return nil
}

func isNearPin(p *image.Point, x, y int) bool {
	return x >= p.X-pinSize/2 && x <= p.X+pinSize/2 &&
		y >= p.Y-pinSize/2 && y <= p.Y+pinSize/2
}

func (b *board) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	for _, r := range b.ropes {
		drawRope(screen, r)
	}

	for _, p := range b.pins {
		drawPin(screen, p)
	}
}

func (b *board) updateRopes() {
	for i := 0; i < len(b.ropes); i++ {
		b.ropes[i].clr = colorGreen
	}

	for i := 0; i < len(b.ropes)-1; i++ {
		for j := i + 1; j < len(b.ropes); j++ {
			ri := &b.ropes[i]
			rj := &b.ropes[j]
			if intersect(*ri.a, *ri.b, *rj.a, *rj.b) {
				ri.clr = colorRed
				rj.clr = colorRed
			}
		}
	}
}
