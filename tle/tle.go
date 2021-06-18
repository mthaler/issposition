package tle

import (
	"bufio"
	"errors"
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

func ReadTLE(scanner bufio.Scanner) (TLE, error) {
	// read the name
	tle := TLE{}
	if (scanner.Scan()) {
		tle.Name = scanner.Text()
	} else {
		return tle, io.EOF
	}

	if (scanner.Scan()) {
		tle.Line1 = scanner.Text()
	} else {
		return tle, errors.New("Line 1 missing")
	}

	if (scanner.Scan()) {
		tle.Line1 = scanner.Text()
	} else {
		return tle, errors.New("Line 2 missing")
	}
	return tle, nil
}