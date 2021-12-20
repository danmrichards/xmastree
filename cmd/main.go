package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"log"
	"math/rand"
	"os"

	tdraw "github.com/danmrichards/xmastree/internal/draw"
	trand "github.com/danmrichards/xmastree/internal/rand"
)

const (
	w, h = 800, 600
	out  = "tree.gif"
)

var (
	// Colours.
	treeGreen = color.RGBA{G: 153, B: 51, A: 255}
	potRed    = color.RGBA{R: 153, A: 255}
	lightBlue = color.RGBA{G: 204, B: 255, A: 255}
	lightRed  = color.RGBA{R: 255, A: 255}
	lightPink = color.RGBA{R: 255, B: 255, A: 255}

	treeLights = []color.RGBA{lightBlue, lightRed, lightPink}
)

func tree(w, h int) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, w, h), palette.Plan9)
	img.SetColorIndex(w/2, h/2, 1)

	// Draw the tree.
	tdraw.TriangleFlatBottom(
		img,
		tdraw.Vertex{
			X: w / 2,
			Y: 30,
		},
		tdraw.Vertex{
			X: w / 4,
			Y: h - (h / 4),
		},
		tdraw.Vertex{
			X: (w / 2) + (w / 4),
			Y: h - (h / 4),
		},
		treeGreen,
		func(img draw.Image, x1, y, x2 int, c color.RGBA) {
			// Fill the tree.
			tdraw.HLine(img, x1, y, x2, c)
		},
	)

	// Draw lights.
	//
	// Use the triangle function again to ensure that we're drawing the lights
	// roughly within the boundaries of the tree, but the callback just adds
	// the lights instead of filling the tree.
	var lastY int
	tdraw.TriangleFlatBottom(
		img,
		tdraw.Vertex{
			X: w / 2,
			Y: 30,
		},
		tdraw.Vertex{
			X: w / 4,
			Y: h - (h / 4),
		},
		tdraw.Vertex{
			X: (w / 2) + (w / 4),
			Y: h - (h / 4),
		},
		treeGreen,
		func(img draw.Image, x1, y, x2 int, c color.RGBA) {
			if (y - lastY) < 40 {
				return
			}
			lastY = y
			tdraw.FilledCircle(img, trand.IntRange(x1, x2), y, 10, treeLights[rand.Intn(len(treeLights))])
			tdraw.FilledCircle(img, trand.IntRange(x1, x2), y, 10, treeLights[rand.Intn(len(treeLights))])
			tdraw.FilledCircle(img, trand.IntRange(x1, x2), y, 10, treeLights[rand.Intn(len(treeLights))])
		},
	)

	// Draw the pot.
	tdraw.Rectangle(
		img,
		tdraw.Vertex{X: (w / 2) - (w / 10), Y: h - (h / 4)},
		tdraw.Vertex{X: (w / 2) + (w / 10), Y: h - (h / 10)},
		potRed,
	)

	return img
}

func main() {
	fmt.Println("Initiating holiday cheer...")

	f, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println("Reticulating splines...")

	// Render 5 frames with a 1-second delay.
	var anim gif.GIF
	for i := 0; i < 5; i++ {
		anim.Image = append(anim.Image, tree(w, h))
		anim.Delay = append(anim.Delay, 100)
	}

	if err = gif.EncodeAll(f, &anim); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Ho! Ho! Ho! Open %s for an xmas treat\n", out)
}
