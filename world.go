package main

import (
	"github.com/banthar/gl"
	"math/rand"
	"time"
)

type Entity struct {
	X, Y int
	*gl.Texture
}

var Player Entity

type Tile struct {
	Blocks bool
	Tex    *gl.Texture
}

var Grass Tile
var Rock Tile

type world struct {
	Grid []Tile
}

var World = new(world)

func InitWorld() {
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
