package orbit

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"log"
	"os"
)

func CreateImage() {
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

	dc := gg.NewContext(1600, 800)
	dc.DrawImage(img, 0, 0)
	dc.DrawCircle(800.0, 400.0, 20.0)
	dc.SetRGB(1.0, 1.0, 1.0)
	dc.Fill()
	err = gg.SaveJPG("images/result.jpg", dc.Image(), 90)
	if err != nil {
		log.Fatal(err)
	}
}