package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	cycles  = 5     // Количество полных колебаний по x
	res     = 0.001 //Угловое разрешение
	size    = 200   // Размер канвы изображения [-size..+size]
	nframes = 64    // Количество фреймов
	delay   = 8     // Задержка между кадрами (в *10мс)
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cycles, _ = strconv.Atoi(r.URL.Query().Get("cycles"))
		res, _ = strconv.ParseFloat(r.URL.Query().Get("res"), 64)
		size, _ = strconv.Atoi(r.URL.Query().Get("size"))
		nframes, _ = strconv.Atoi(r.URL.Query().Get("nframes"))
		delay, _ = strconv.Atoi(r.URL.Query().Get("delay"))

		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func lissajous(out io.Writer) {
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	palette := make([]color.Color, 0, nframes)
	palette = append(palette, color.RGBA{A: 255})
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	r := float64(rand.Intn(200))
	g := float64(rand.Intn(200))
	b := float64(rand.Intn(200))
	for i := 0; i < nframes; i++ {
		scale := float64(i) / float64(nframes)
		c := color.RGBA{R: uint8(55 + r*scale), G: uint8(55 + g*scale), B: uint8(55 + b*scale), A: 255}
		palette = append(palette, c)
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), uint8((i%(len(palette)-1))+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	err := gif.EncodeAll(out, &anim)
	if err != nil {
		log.Fatal(err)
	}
}
