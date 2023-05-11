package main

import (
	"fmt"
	"net/http"

	"github.com/youzhicode/golang/ch1/lissajous/lissajous"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "have error %v\n", err)
		}
		lissajous.Lissajous(w)
		/*if cycles, err := strconv.Atoi(r.FormValue("cycles")); err != nil {
			fmt.Fprintf(w, "500 error %v", err)
		} else {
			//lissajous(w, float64(cycles))
		}*/
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe("Localhost:8080", nil)
}

/*func lissajous(out io.Writer, cycles float64) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	if cycles == 0 {
		cycles = 5
	}

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
}*/
