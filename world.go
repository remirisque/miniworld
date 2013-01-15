package main

import (
	"github.com/banthar/gl"
	"math/rand"
	"time"
)

type Entity struct {
	x, y, mag int //starting location and magnitude (speed in "tiles per second") of entity
	moving    int
	moveTick  time.Time // set when entity moves
	texture   *gl.Texture
}

var player Entity

type Tile struct {
	blocks  bool
	texture *gl.Texture
}

var grass Tile
var rock Tile

type World struct {
	grid []Tile
}

var world = new(World)

func (e *Entity) update() {
	if e.moving != IDLE {
		player.move()
	}
}
func initWorld() {
	grass = Tile{false, grassTexture}
	rock = Tile{true, rockTexture}
	player = Entity{5, 5, 10, IDLE, time.Now(), playerTexture}
	rand.Seed(int64(time.Now().Nanosecond()))
	world.grid = make([]Tile, GridHeight*GridWidth)
	for idx := 0; idx < len(world.grid); idx++ {
		if (idx < GridWidth) || (idx > GridWidth*(GridHeight-1)-1) { //walls of rock at top and bottom edge
			world.grid[idx] = rock
		} else if (idx%GridWidth == 0) || (idx%GridWidth == GridWidth-1) { //walls of rock at left and right edge
			world.grid[idx] = rock
		} else if rand.Float64() < .1 {
			world.grid[idx] = rock
		} else {
			world.grid[idx] = grass
		}
	}

}
