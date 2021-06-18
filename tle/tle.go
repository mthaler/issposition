package tle

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

type TLE struct {
	Name string
	Line1 string
	Line2 string
}

func ReadTLEs(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
