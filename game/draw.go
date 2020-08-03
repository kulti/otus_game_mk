package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	colorRed   = color.RGBA{255, 0, 0, 255}
	colorGreen = color.RGBA{0, 255, 0, 255}
)

const pinSize = 10

func drawRope(screen *ebiten.Image, rope rope) {
	ebitenutil.DrawLine(screen, float64(rope.a.X), float64(rope.a.Y), float64(rope.b.X), float64(rope.b.Y), rope.clr)
}

func drawPin(screen *ebiten.Image, pin *image.Point) {
	ebitenutil.DrawRect(screen, float64(pin.X-pinSize/2), float64(pin.Y-pinSize/2), pinSize, pinSize, color.Black)
}
