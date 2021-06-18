package main

import (
	"bufio"
	"fmt"
	"github.com/mthaler/iss-position/download"
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

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}