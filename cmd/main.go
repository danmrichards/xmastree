package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/danmrichards/xmastree/internal/draw"
)

const w, h = 800, 600

var (
	img   = image.NewRGBA(image.Rect(0, 0, w, h))
	green = color.RGBA{G: 153, B: 51, A: 255}
	red   = color.RGBA{R: 200, A: 255}
)

func main() {
	// Draw the tree.
	draw.TriangleFlatBottom(
		img,
		draw.Vertex{
			X: w / 2,
			Y: 30,
		},
		draw.Vertex{
			X: w / 4,
			Y: h - (h / 4),
		},
		draw.Vertex{
			X: (w / 2) + (w / 4),
			Y: h - (h / 4),
		},
		green,
	)

	// Draw the pot.
	draw.Rectangle(
		img,
		draw.Vertex{X: (w / 2) - (w / 10), Y: h - (h / 4)},
		draw.Vertex{X: (w / 2) + (w / 10), Y: h - (h / 10)},
		red,
	)

	f, err := os.Create("draw.png")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err = png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}
