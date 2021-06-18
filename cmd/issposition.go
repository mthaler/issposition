package main

import (
	"fmt"
	"github.com/joshuaferrara/go-satellite"
	"github.com/mthaler/iss-position/internal/download"
	"github.com/mthaler/iss-position/internal/orbit"
	"github.com/mthaler/iss-position/internal/tle"
	"log"
	"net/http"
	"os"
)

func main() {
	fileUrl := "https://www.celestrak.com/NORAD/elements/stations.txt"
	filePath := "stations.txt"
	err := download.DownloadFile(filePath, fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tles, err := tle.ReadTLEs(file)
	if err != nil {
		panic(err)
	}

	tle, ok := tles["ISS (ZARYA)"]
	if !ok {
		panic("ISS TLE not found")
	}

	iss := satellite.TLEToSat(tle.Line1, tle.Line2, "wgs84")
	orbit.CreateImage(iss)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}