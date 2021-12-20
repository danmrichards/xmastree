package draw

import (
	"image"
	"image/color"
	"image/draw"
)

// Vertex is a point in 2D space.
type Vertex struct {
	X, Y int
}

// TriangleScanLineFunc is a function that is run for each scanline when drawing
// a triangle.
type TriangleScanLineFunc func(img draw.Image, x1, y, x2 int, c color.RGBA)

// HLine draws a horizontal line
func HLine(img draw.Image, x1, y, x2 int, c color.RGBA) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, c)
	}
}

// TriangleFlatBottom draws a flat-bottom triangle, defined by the 3 vertices,
// filled with the given colour.
func TriangleFlatBottom(img draw.Image, v1, v2, v3 Vertex, c color.RGBA, f TriangleScanLineFunc) {
	invslope1 := float32(v2.X-v1.X) / float32(v2.Y-v1.Y)
	invslope2 := float32(v3.X-v1.X) / float32(v3.Y-v1.Y)

	curx1 := float32(v1.X)
	curx2 := float32(v1.X)

	for scanlineY := v1.Y; scanlineY <= v2.Y; scanlineY++ {
		f(img, int(curx1), scanlineY, int(curx2), c)
		curx1 += invslope1
		curx2 += invslope2
	}
}

// Rectangle draws a rectangle, defined by the two vertices (upper left and
// lower right), filled with the given colour.
func Rectangle(img draw.Image, v1, v2 Vertex, c color.RGBA) {
	draw.Draw(
		img,
		image.Rect(v1.X, v1.Y, v2.X, v2.Y),
		&image.Uniform{C: c},
		image.Point{},
		draw.Src,
	)
}

// Circle draws a circle with the given radius and colour.
func Circle(img draw.Image, x0, y0, r int, c color.Color) {
	x, y, dx, dy := r-1, 0, 1, 1
	err := dx - (r * 2)

	for x > y {
		img.Set(x0+x, y0+y, c)
		img.Set(x0+y, y0+x, c)
		img.Set(x0-y, y0+x, c)
		img.Set(x0-x, y0+y, c)
		img.Set(x0-x, y0-y, c)
		img.Set(x0-y, y0-x, c)
		img.Set(x0+y, y0-x, c)
		img.Set(x0+x, y0-y, c)

		if err <= 0 {
			y++
			err += dy
			dy += 2
		}
		if err > 0 {
			x--
			dx += 2
			err += dx - (r * 2)
		}
	}
}

func FilledCircle(img draw.Image, x0, y0, r int, c color.Color) {
	for dr := r; dr > 0; dr-- {
		Circle(img, x0, y0, dr, c)
	}
}
