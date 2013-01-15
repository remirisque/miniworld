package main

import (
	"github.com/banthar/gl"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"log"
	"os"
)

var grassTexture *gl.Texture
var rockTexture *gl.Texture
var playerTexture *gl.Texture

func loadTexture(path string) (*gl.Texture, error) {
	file, err := os.Open(path)
	if err != nil {

	}
	defer file.Close()

	srcimg, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	// convert to RGBA if not already
	b := srcimg.Bounds()
	var dstimg *image.RGBA
	if srcimg.ColorModel() != color.RGBAModel {
		dstimg = image.NewRGBA(b)
		draw.Draw(dstimg, b, srcimg, b.Min, draw.Src)
	} else {
		dstimg = srcimg.(*image.RGBA)
	}

	tex := gl.GenTexture()
	tex.Bind(gl.TEXTURE_2D)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, b.Dx(), b.Dy(), 0, gl.RGBA, gl.UNSIGNED_BYTE, dstimg.Pix)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	return &tex, nil
}

func coordToOffset(x, y int) int {
	return y*GridWidth + x
}

func offsetToCoord(offset int) (int, int) {
	return offset % GridWidth, offset / GridWidth
}

func renderTile(x, y float32, tex *gl.Texture) {
	tex.Bind(gl.TEXTURE_2D)
	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0, 0)
	gl.Vertex2f(x, y)
	gl.TexCoord2f(1, 0)
	gl.Vertex2f(x+1, y)
	gl.TexCoord2f(1, 1)
	gl.Vertex2f(x+1, y+1)
	gl.TexCoord2f(0, 1)
	gl.Vertex2f(x, y+1)
	gl.End()
}

func initResources() {
	var err error
	grassTexture, err = loadTexture("res/grass.png")
	rockTexture, err = loadTexture("res/rock.png")
	playerTexture, err = loadTexture("res/char.png")
	if err != nil {
		log.Fatal(err)
	}

}
