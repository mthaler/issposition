package main

import (
	"bytes"
	"github.com/joshuaferrara/go-satellite"
	"github.com/mthaler/iss-position/internal/download"
	"github.com/mthaler/iss-position/internal/orbit"
	"github.com/mthaler/iss-position/internal/tle"
	"image/jpeg"
	"log"
	"net/http"
)

func main() {
	fileUrl := "https://www.celestrak.com/NORAD/elements/stations.txt"
	buf, err := download.Download(fileUrl)
	if err != nil {
		panic(err)
	}

	tles, err := tle.ReadTLEs(buf)
	if err != nil {
		panic(err)
	}

	tle, ok := tles["ISS (ZARYA)"]
	if !ok {
		panic("ISS TLE not found")
	}

	iss := satellite.TLEToSat(tle.Line1, tle.Line2, "wgs84")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/map", mapHandler(iss))

	log.Println("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func mapHandler(iss satellite.Satellite) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		img := orbit.CreateImage(iss)
		buf := new(bytes.Buffer)
		var opt jpeg.Options
		opt.Quality = 90.0
		err := jpeg.Encode(buf, img, &opt)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(buf.Bytes())
	}
}