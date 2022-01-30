// server1 is a minimal web server
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 255}, color.RGBA{0, 0, 255, 255},
	color.RGBA{0, 150, 100, 255}, color.RGBA{175, 100, 0, 255}}

// settings contains all of the GIF settings. Using interface as type to support float64 and int.
var settings = make(map[string]interface{})

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lissajous(w, r)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func lissajous(out io.Writer, r *http.Request) {
	params := r.URL.Query()
	defaults(params)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: settings["nframes"].(int)}
	phase := 0.0 // phase difference
	for i := 0; i < settings["nframes"].(int); i++ {
		rect := image.Rect(0, 0, 2*settings["size"].(int)+1, 2*settings["size"].(int)+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(settings["cycles"].(int))*2*math.Pi; t += settings["res"].(float64) {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(settings["size"].(int)+int(x*float64(settings["size"].(int))+0.5),
				settings["size"].(int)+int(y*float64(settings["size"].(int))+0.5), uint8(len(palette)-rand.Intn(len(palette)-1)))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, settings["delay"].(int))
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

// defaults loads the default values for settings. it then iterates over settings,
// loading values set in the URL parameters, if they exist.
func defaults(params url.Values) {
	settings["cycles"] = 5
	settings["size"] = 100
	settings["res"] = 0.001
	settings["nframes"] = 64
	settings["delay"] = 8
	for key, value := range settings {
		if _, ok := params[key]; ok {
			switch value.(type) {
			case int:
				settings[key], _ = strconv.Atoi(params[key][0])
			case float64:
				settings[key], _ = strconv.ParseFloat(params[key][0], 32)
			}
		}
	}
}
