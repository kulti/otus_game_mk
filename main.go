package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/kulti/otus_game_mk/game"
)

func main() {
	ebiten.SetWindowSize(600, 600)

	ebiten.RunGame(game.New())
}
