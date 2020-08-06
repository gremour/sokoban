package game

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	Wall         = iota
	Floor        = iota
	Target       = iota
	Box          = iota
	DeliveredBox = iota
	Player       = iota
)

var mapObjects = map[rune]byte{
	'.': Target,
	'#': Wall,
	'o': Box,
	'0': DeliveredBox,
	'@': Player,
}

// Map ...
type Map struct {
	Level  []byte
	Width  int
	Height int
}

// Pos ...
type Pos struct {
	X, Y int
}

// GameToIndex ...
func (m Map) GameToIndex(x, y int) int {
	return x + y*m.Width
}

// IndexToGame ...
func (m Map) IndexToGame(ind int) (int, int) {
	x := ind % m.Width
	y := ind / m.Width
	return x, y
}

// IsInBounds ...
func (m Map) IsInBounds(x, y int) bool {
	return x >= 0 && x < m.Width && y >= 0 && y < m.Height
}

// ObjAt ...
func (m Map) ObjAt(x, y int) byte {
	if !m.IsInBounds(x, y) {
		return Wall
	}
	return m.Level[m.GameToIndex(x, y)]
}

// PutObjAt ...
func (m Map) PutObjAt(x, y int, obj byte) {
	if !m.IsInBounds(x, y) {
		return
	}
	m.Level[m.GameToIndex(x, y)] = obj
}

// MapFromFile ...
func MapFromFile(filename string) (Map, Pos, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Map{}, Pos{}, err
	}

	level := strings.ReplaceAll(string(bytes), "\r", "")
	lines := strings.Split(level, "\n")

	m := Map{}
	var player *Pos
	boxes := 0
	targets := 0
	for i, l := range lines {
		if m.Width != 0 && len(l) != m.Width && len(l) == 0 {
			return Map{}, Pos{}, fmt.Errorf("invalid map: line %v lenght (%v) is different from %v", i, len(l), m.Width)
		}
		if m.Width == 0 {
			m.Width = len(l)
		}
		for j, r := range l {
			obj, ok := mapObjects[r]
			if !ok {
				obj = Floor
			}
			if obj == Player {
				player = &Pos{
					X: j,
					Y: i,
				}
				obj = Floor
			}
			if obj == Box || obj == DeliveredBox {
				boxes++
			}
			if obj == Target || obj == DeliveredBox {
				targets++
			}
			m.Level = append(m.Level, obj)
		}
	}
	if boxes == 0 || targets == 0 {
		return Map{}, Pos{}, fmt.Errorf("need at least one box and target")
	}
	if boxes != targets {
		return Map{}, Pos{}, fmt.Errorf("boxes and targets do not match")
	}
	if player == nil {
		return Map{}, Pos{}, fmt.Errorf("no player in map")
	}
	m.Height = len(lines)
	return m, *player, nil
}
