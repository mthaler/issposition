package main

import (
	"fmt"
	"github.com/mthaler/iss-position/download"
)

func main() {
	fileUrl := "https://www.celestrak.com/NORAD/elements/stations.txt"
	err := download.DownloadFile("logo.svg", fileUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + fileUrl)
}