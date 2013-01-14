package main

import (
	"github.com/banthar/gl"
	"github.com/jteeuwen/glfw"
	"time"
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
	glfw.SetSwapInterval(true)
	glfw.SetKeyCallback(inputCallback)

	gl.Ortho(0, 40, 0, 30, 0, 0)
	for Running {
		if (1 / time.Since(dt) / time.Millisecond) > 15 {
			dt := time.Now()
			util.Render()
			glfw.SwapBuffers()
		}
	}
}

func inputCallback(key int, state int) {
	if key == glfw.KeyEsc && state == glfw.KeyPress {
		running = false
	}
}
