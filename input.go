package main

import (
	"github.com/jteeuwen/glfw"
	"time"
)

const (
	UP    = -GridWidth
	DOWN  = GridWidth
	LEFT  = -1
	RIGHT = 1
	IDLE  = 0
)

func inputCallback(key int, state int) {
	if key == glfw.KeyEsc && state == glfw.KeyPress {
		Running = false
	} else if key == glfw.KeyUp && state == glfw.KeyPress {
		player.Moving(UP)
	} else if key == glfw.KeyDown && state == glfw.KeyPress {
		player.Moving(DOWN)
	} else if key == glfw.KeyLeft && state == glfw.KeyPress {
		player.Moving(LEFT)
	} else if key == glfw.KeyRight && state == glfw.KeyPress {
		player.Moving(RIGHT)
	} else if (key == glfw.KeyUp || key == glfw.KeyDown || key == glfw.KeyLeft || key == glfw.KeyRight) && state == glfw.KeyRelease {
		player.Moving(IDLE)
	}
}

func (e *Entity) move() {
	targetOffset := coordToOffset(e.x, e.y) + e.moving //get offset of tile entity is moving towards
	tx, ty := offsetToCoord(targetOffset)              //and get the coordinate pair of target as well
	//verify no obstacles first
	if world.grid[targetOffset].blocks {
		return
	}

	if time.Since(e.moveTick).Seconds() > 1.0/float64(e.mag) {
		e.x, e.y = tx, ty
		e.moveTick = time.Now()
	}
}

func (e *Entity) Moving(dir int) {
	e.moving = dir
	e.moveTick = time.Now()
}
