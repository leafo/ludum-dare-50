package loops

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"io"

	gl "github.com/go-gl/gl/v3.3-core/gl"
)

type Texture struct {
	Texture uint32
}

func NewTexture(r io.Reader) Texture {
	img, err := png.Decode(r)

	rgbaImg := image.NewRGBA(img.Bounds())

	draw.Draw(rgbaImg, rgbaImg.Bounds(), img, image.Pt(0, 0), draw.Src)

	if err != nil {
		panic(err)
	}

	texture := Texture{0}

	gl.GenTextures(1, &texture.Texture)
	fmt.Println("Gen texture id:", texture.Texture)

	gl.ActiveTexture(gl.TEXTURE0)
	texture.Bind()

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_R, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.SRGB_ALPHA, int32(rgbaImg.Rect.Size().X), int32(rgbaImg.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgbaImg.Pix))

	// gl.TexImage2D(GL_TEXTURE_2D, 0, GL_RGBA, w, h, 0, format, type, bytes)

	return texture
}

func (self *Texture) Bind() {
	gl.BindTexture(gl.TEXTURE_2D, self.Texture)
}
