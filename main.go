package main

import (
	"bytes"
	_ "embed"
	"image/color"
	"runtime"

	"github.com/leafo/ludum-dare-50/loops"
)

//go:embed hi.png
var image []byte

func init() {
	runtime.LockOSThread()
}

func main() {
	loop := loops.NewLoopWindow()
	loop.Title = "Ludum Dare 50"

	var texture loops.Texture

	loop.Load = func() {
		texture = loops.NewTexture(bytes.NewReader(image))
	}

	loop.Draw = func(t float64, g *loops.Graphics) {
		g.SetColor(color.RGBA{20, 20, 20, 255})
		g.SetMat(loops.NewIdentityMat4())
		g.DrawRect(-1, -1, 2, 2)

		g.SetColor(color.RGBA{255, 20, 20, 255})
		g.DrawRect(-0.5, -0.5, 1, 1)

		texture.Bind()
	}

	loop.Run()
}
