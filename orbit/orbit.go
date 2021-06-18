package orbit

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/joshuaferrara/go-satellite"
	"image"
	"log"
	"math"
	"os"
)

const w = 1600
const h = 800

func CreateImage(pos satellite.LatLong) {
	dc := gg.NewContext(w, h)
	drawMap(dc)
	drawISS(dc, pos)
	err := gg.SaveJPG("images/result.jpg", dc.Image(), 90)
	if err != nil {
		log.Fatal(err)
	}
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

func drawISS(dc *gg.Context, pos satellite.LatLong) {
	x := w / 2.0 + pos.Longitude * w / (2.0 * math.Pi)
	y := h / 2.0 - pos.Latitude * h / math.Pi
	dc.DrawCircle(x, y, 15.0)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()
}