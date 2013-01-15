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

var GrassTexture *gl.Texture
var RockTexture *gl.Texture

func LoadTexture(path string) (*gl.Texture, error) {
	file, err := os.Open(path)
	if err != nil {

	}
	defer file.Close()

	srcimg, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	// convert to NRGBA if not already
	b := srcimg.Bounds()
	var dstimg *image.NRGBA
	if srcimg.ColorModel() != color.NRGBAModel {
		dstimg = image.NewNRGBA(b)
		draw.Draw(dstimg, b, srcimg, b.Min, draw.Src)
	} else {
		dstimg = srcimg.(*image.NRGBA)
	}

	tex := gl.GenTexture()
	tex.Bind(gl.TEXTURE_2D)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, b.Dx(), b.Dy(), 0, gl.RGBA, gl.UNSIGNED_BYTE, dstimg.Pix)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	return &tex, nil
}

func CoordToOffset(x, y int) int {
	return y*GridWidth + x
}

func OffsetToCoord(offset int) (int, int) {
	return offset % GridWidth, offset / GridWidth
}

func RenderTile(x, y float32, tex *gl.Texture) {
	tex.Bind(gl.TEXTURE_2D)
	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0, 0)
	gl.Vertex2f(x, y)
	gl.TexCoord2f(0+1, 0)
	gl.Vertex2f(x+1, y)
	gl.TexCoord2f(0+1, 0+1)
	gl.Vertex2f(x+1, y+1)
	gl.TexCoord2f(0, 0+1)
	gl.Vertex2f(x, y+1)
	gl.End()
}

func InitResources() {
	var err error
	GrassTexture, err = LoadTexture("res/grass.png")
	RockTexture, err = LoadTexture("res/rock.png")
	if err != nil {
		log.Fatal(err)
	}

}
