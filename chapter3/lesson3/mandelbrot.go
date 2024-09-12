package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"math/rand"
	"os"
	"time"
)

var palettes []color.RGBA

func generateRndPalette(n int) []color.RGBA {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	colors := make([]color.RGBA, n)
	for i := range colors {
		colors[i] = color.RGBA{
			R: uint8(rand.Intn(256)),
			G: uint8(rand.Intn(256)),
			B: uint8(rand.Intn(256)),
			A: 255,
		}
	}
	return colors

}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	palettes = generateRndPalette(100)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	f, err := os.Create("mandelbrot.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := uint16(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return palettes[n%16]
		}
	}
	return color.Black
}
