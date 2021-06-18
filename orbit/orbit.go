package orbit

import (
	"fmt"
	"github.com/fogleman/gg"
	"github.com/joshuaferrara/go-satellite"
	"image"
	"log"
	"os"
)

func CreateImage(pos satellite.LatLong) {
	dc := gg.NewContext(1600, 800)
	drawMap(dc)
	drawISS(dc)
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

func drawISS(dc *gg.Context) {
	dc.DrawCircle(800.0, 400.0, 20.0)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()
}