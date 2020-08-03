package game

import (
	"github.com/hajimehoshi/ebiten"
)

type Game struct {
	board *board
}

func New() *Game {
	return &Game{
		board: newBoard(),
	}
}

func (g *Game) Update(_ *ebiten.Image) error {
	g.board.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 400, 400
}
