package main

import (
	"github.com/banthar/gl"
	"github.com/jteeuwen/glfw"
	"time"
)

const (
	GridHeight = 15
	GridWidth  = 20
)

var Running = true
var DT = time.Now()

func main() {
	glfw.Init()
	defer glfw.Terminate()

	glfw.OpenWindow(640, 480, 8, 8, 8, 8, 0, 0, glfw.Windowed)
	defer glfw.CloseWindow()

	glfw.SetWindowTitle("Tile test")
	glfw.Enable(glfw.StickyKeys)
	glfw.SetSwapInterval(1)
	glfw.SetKeyCallback(inputCallback)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Ortho(0, GridWidth, GridHeight, 0, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE_2D)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	initResources()
	initWorld()

	for Running {
		if (time.Since(DT).Nanoseconds() / 1000000) > 15 { //don't loop faster than every 15ms
			DT = time.Now()
			gl.Clear(gl.COLOR_BUFFER_BIT)
			player.update()
			renderScene()
			glfw.SwapBuffers()
		}
	}
}

func renderScene() {
	for idx := 0; idx < len(world.grid); idx++ {
		x, y := offsetToCoord(idx)
		renderTile(float32(x), float32(y), world.grid[idx].texture)
	}
	//render player
	renderTile(float32(player.x), float32(player.y), player.texture)
}
