package world

import (
	"github.com/banthar/gl"
	"math/rand"
	"time"
	"util"
)

const (
	GridWidth  = 40
	GridHeight = 30
)

type Entity struct {
	X, Y int
	*gl.Texture
}

var Player Entity

type Tile struct {
	Blocks  bool
	Texture *gl.Texture
}

var Grass = new(Tile{false, util.GrassTexture})
var Rock = new(Tile{true, util.RockTexture})

type world struct {
	Grid []Tile
}

var World = new(world)

func init() {
	rand.Seed(time.Now())
	World.Grid = make([]Tile, 1200)
	for idx := 0; idx < len(World.Grid); idx++ {
		if rand.Float64() < .1 {
			World.Grid[idx] = new(Rock)
		} else {
			World.Grid[idx] = new(Grass)
		}
	}
}
