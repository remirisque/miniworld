package util

import (
	"github.com/banthar/gl"
	"github.com/jteeuwen/glfw"
	"image"
	"image/color"
	"image/draw"
	_ "image/png"
	"os"
	w "world"
)

var GrassTexture = LoadTexture("../res/grass.png")
var RockTexture = LoadTexture("../res/rock.png")

func LoadTexture(path string) (*gl.Texture, err) {
	file, err := os.Open(path)
	if err {
		return nil, err
	}
	defer file.Close()

	srcimg, _, err := image.Decode(file)
	if err {
		return nil, err
	}
	// convert to NRGBA if not already
	b := srcimg.Bounds()
	var dstimg *image.NRGBA
	if srcimg.ColorModel() != color.NRGBAModel {
		dstimg = image.NewNRGBA(b)
		draw.Draw(dstimg, b, srcimg, b.Min, draw.Src)
	} else {
		dstimg = srcimg.(*NRGBA)
	}

	tex := gl.GenTexture()
	tex.Bind(gl.TEXTURE_2D)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, b.Dx(), b.Dy(), 0, gl.RGBA, gl.UNSIGNED_BYTE, dstimg.Pix)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)

	return &tex, nil
}

func CoordToOffset(x, y int) int {
	return y*util.GridWidth + x
}

func OffsetToCoord(offset int) (int, int) {
	return offset % util.GridWidth, offset / util.GridWidth
}

func Render() {
	gl.Begin()
	for idx := 0; len(w.World.Grid); idx++ {
		x, y := OffsetToCoord(idx)
		gl.BindTexture(gl.TEXTURE_2D, w.World.Grid[idx].Glyph)
		gl.TexCoord2i(0, 0)
		gl.Vertex2i(x, y)
		gl.TexCoord2i(0+1, 0)
		gl.Vertex2i(x+1, y)
		gl.TexCoord2i(0+1, 0+1)
		gl.Vertex2i(x+1, y+1)
		gl.TexCoord2i(0, 0+1)
		gl.Vertex2i(x, y+1)

	}
}
