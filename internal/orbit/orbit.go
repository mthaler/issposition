package orbit

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/joshuaferrara/go-satellite"
	"image"
	"log"
	"math"
	"os"
	"time"
)

const w = 1600
const h = 800

func CreateImage(iss satellite.Satellite) {
	dc := gg.NewContext(w, h)
	drawMap(dc)
	drawLatLngGrid(dc)
	drawOrbit(dc, iss)
	drawISS(dc, iss)
	err := gg.SaveJPG("images/result.jpg", dc.Image(), 90)
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

func toScreen(latLng satellite.LatLong) (float64, float64) {
	x := w / 2.0 + latLng.Longitude * w / (2.0 * math.Pi)
	if x < 0.0 {
		x = x + w
	} else if x > w {
		x = x - w
	}
	y := h / 2.0 - latLng.Latitude * h / math.Pi
	return x, y
}

func drawMap(dc *gg.Context) {
	file, err := os.Open("images/map.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, fmtName, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Format:", fmtName)
	dc.DrawImage(img, 0, 0)
}

func drawLatLngGrid(dc *gg.Context) {
	dc.SetRGBA(1.0, 1.0, 1.0, 0.2)
	dc.SetLineWidth(3)
	dc.MoveTo(0, h / 2.0)
	dc.LineTo(w, h / 2.0)
	dc.MoveTo(w / 4.0, 0)
	dc.LineTo(w / 4.0, h)
	dc.MoveTo(w / 2.0, 0)
	dc.LineTo(w / 2.0, h)
	dc.MoveTo(w / 4.0 * 3.0, 0)
	dc.LineTo(w / 4.0 * 3.0, h)
	dc.Stroke()
}

func drawOrbit(dc *gg.Context, iss satellite.Satellite) {
	now := time.Now().UTC()
	start :=  now.Add(-time.Minute * 60)
	end := now.Add(time.Minute * 60)
	t := start
	doDraw := false
	dc.SetRGBA(1.0, 1.0, 1.0, 0.6)
	dc.SetLineWidth(4)
	var previousX float64
	for t.Before(end) {
		pos, _ := propagate(iss, t)
		gmst := gsTimeFromDate(t)
		_, _, latLng := satellite.ECIToLLA(pos, gmst)
		x, y := toScreen(latLng)
		if doDraw {
			if math.Abs(x - previousX) < w / 2.0 {
				dc.LineTo(x, y)
			} else {
				dc.MoveTo(x, y)
			}
		} else {
			dc.MoveTo(x, y)
		}
		doDraw = true
		previousX = x
		t = t.Add(time.Minute)
	}
	dc.Stroke()
}

func drawISS(dc *gg.Context, iss satellite.Satellite) {
	now := time.Now().UTC()
	pos, _ := propagate(iss, now)
	gmst := gsTimeFromDate(now)
	_, _, latLng := satellite.ECIToLLA(pos, gmst)
	x, y := toScreen(latLng)
	dc.DrawCircle(x, y, 15.0)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()
}