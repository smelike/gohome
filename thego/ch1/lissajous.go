package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"net/http"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:9000", nil)
}

// 输出到 os.Writer
func lissajous(out io.Writer, r *http.Request) {
	const (
		cycles  = 1
		res     = 0.001
		size    = 400
		nframes = 2
		delay   = 200
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)

	}
	gif.EncodeAll(out, &anim)
}

/*
Exercise 1.5: Change the Lissajous program's color palette to green on block, for added
authenticity. To create the web color #RRGGBB, use color.RGBA{0xRR, 0xGG, 0xBB,0xff},
where each pair of hexadecimal digits represents the intencity of the red, green,
or blue component of the pixel.

Exercise 1.6: Modify the Lissajous program to produce images in multiple colos by adding
more values to palette and then displaying them by changing the third argument of
SetColorIndex in some interesting way.

Exercise 1.12: Modify the Lissajous server to read parameter values from the URL.
For example, you might arrange it so that a URL like http://localhost:8000/?cycle=20 sets
the number of cycles to 20 instead of the default 5. Use the strconv.Atoi function to convert
the string parameter into an integer. You can see its documentation with go doc strconv.Atoi.

*/
