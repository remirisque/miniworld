package main

import (
	"github.com/banthar/gl"
	"github.com/jteeuwen/glfw"
	"time"
)

const (
	GridHeight = 30
	GridWidth  = 40
)

var Running = true

func main() {
	dt := time.Now()

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
	gl.Ortho(0, 40, 0, 30, -1, 1)
	gl.MatrixMode(gl.MODELVIEW)
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE_2D)
	gl.ClearColor(0.0, 0.0, 0.0, 1.0)
	InitResources()
	InitWorld()

	for Running {
		if (time.Since(dt).Nanoseconds() / 1000000) > 15 {
			dt = time.Now()
			gl.Clear(gl.COLOR_BUFFER_BIT)
			renderScene()
			glfw.SwapBuffers()
		}
	}
}

func inputCallback(key int, state int) {
	if key == glfw.KeyEsc && state == glfw.KeyPress {
		Running = false
	}
}

func renderScene() {
	for idx := 0; idx < len(World.Grid); idx++ {
		x, y := OffsetToCoord(idx)
		RenderTile(float32(x), float32(y), World.Grid[idx].Tex)
	}
}
