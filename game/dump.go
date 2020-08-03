// +build !release

package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func (b *board) dumpLevel() {
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		fmt.Println("=== DUMP LEVEL ===")
		for _, r := range b.ropes {
			fmt.Println(r.a.X, r.a.Y, r.b.X, r.b.Y)
		}
	}
}
