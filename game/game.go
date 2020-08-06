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
	Win     bool
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

// Move ...
func (g *Game) Move(dx, dy int) {
	if !g.CanMove(dx, dy) {
		return
	}
	x, y := g.Player.X+dx, g.Player.Y+dy
	obj := g.Map.ObjAt(x, y)
	if obj == Box || obj == DeliveredBox {
		nx, ny := x+dx, y+dy
		nobj := g.Map.ObjAt(nx, ny)
		if nobj == Floor {
			g.Map.PutObjAt(nx, ny, Box)
		} else {
			g.Map.PutObjAt(nx, ny, DeliveredBox)
		}
		if obj == Box {
			g.Map.PutObjAt(x, y, Floor)
		} else {
			g.Map.PutObjAt(x, y, Target)
		}
	}
	g.Player.X += dx
	g.Player.Y += dy
	g.CheckWin()
}

// CanMove ...
func (g *Game) CanMove(dx, dy int) bool {
	x, y := g.Player.X+dx, g.Player.Y+dy
	obj := g.Map.ObjAt(x, y)
	switch obj {
	case Floor:
		fallthrough
	case Target:
		return true
	case Box:
		fallthrough
	case DeliveredBox:
		nx, ny := x+dx, y+dy
		nobj := g.Map.ObjAt(nx, ny)
		if nobj == Floor || nobj == Target {
			return true
		}
	}
	return false
}

// CheckWin ...
func (g *Game) CheckWin() {
	win := true
	for _, o := range g.Map.Level {
		if o == Box {
			win = false
			break
		}
	}
	g.Win = win
}
