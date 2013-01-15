package main

import (
	"github.com/go-gl/glh"
	"math/rand"
	"time"
)

const (
	GridHeight = 30
	GridWidth  = 40
)

type Entity struct {
	X, Y int
	*glh.Texture
}

var Player Entity

type Tile struct {
	Blocks bool
	Tex    *glh.Texture
}

var Grass Tile
var Rock Tile

type world struct {
	Grid []Tile
}

var World = new(world)

func InitHARPDARPTHESECOND() {
	Grass = Tile{false, GrassTexture}
	Rock = Tile{true, RockTexture}
	rand.Seed(int64(time.Now().Nanosecond()))
	World.Grid = make([]Tile, GridHeight*GridWidth)
	for idx := 0; idx < len(World.Grid); idx++ {
		if rand.Float64() < .1 {
			World.Grid[idx] = Rock
		} else {
			World.Grid[idx] = Grass
		}
	}
}
