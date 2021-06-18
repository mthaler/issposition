package main

import (
	"fmt"
	"github.com/mthaler/iss-position/download"
	"github.com/mthaler/iss-position/tle"
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

	tle.ReadTLEs(file)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :3000...")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}