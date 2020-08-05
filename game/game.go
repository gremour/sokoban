package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

// Game ...
type Game struct {
	Map     Map
	Player  Pos
	Sprites []*pixel.Sprite
}

// New ...
func New(levelName string, sprites []*pixel.Sprite) (*Game, error) {
	m, p, err := MapFromFile(levelName)
	if err != nil {
		return nil, err
	}
	return &Game{
		Map:     m,
		Player:  p,
		Sprites: sprites,
	}, nil
}

// Draw ...
func (g *Game) Draw(win *pixelgl.Window, m pixel.Matrix) {
	for i, v := range g.Map.Level {
		x, y := g.Map.IndexToGame(i)
		curm := m.Moved(pixel.V(float64(x)*32, float64(g.Map.Height-y)*32))
		g.Sprites[v].Draw(win, curm)
	}
}

// DrawPlayer ...
func (g *Game) DrawPlayer(win *pixelgl.Window, m pixel.Matrix) {
	curm := m.Moved(pixel.V(float64(g.Player.X)*32, float64(g.Map.Height-g.Player.Y)*32))
	g.Sprites[Player].Draw(win, curm)
}
