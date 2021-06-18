package main

import (
	"fmt"
	"github.com/mthaler/iss-position/download"
	"github.com/mthaler/iss-position/tle"
	"log"
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

	tle.ReadTLEs(file)
}