package main

import (
	"fmt"
	"github.com/joshuaferrara/go-satellite"
	"github.com/mthaler/iss-position/download"
	"github.com/mthaler/iss-position/tle"
	"log"
	"net/http"
	"os"
	"time"
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
	fmt.Println(iss)

	now := time.Now().UTC()
	fmt.Println(now)

	pos, vel := propagate(iss, now)
	fmt.Println("Position: ", pos, "velocity:", vel)

	gmst := gsTimeFromDate(now)
	altitude, velocity, latLng := satellite.ECIToLLA(pos, gmst)
	fmt.Println("altitude: ", altitude, "velocity: ", velocity, "latLng: ", latLng)

	latLngDeg := satellite.LatLongDeg(latLng)
	fmt.Println("latLngDeg:", latLngDeg)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func propagate(sat satellite.Satellite, t time.Time) (position, velocity satellite.Vector3) {
	return satellite.Propagate(sat, t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}

func gsTimeFromDate(t time.Time) float64 {
	return satellite.GSTimeFromDate(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
}