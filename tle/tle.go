package tle

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

type TLE struct {
	Name string
	Line1 string
	Line2 string
}

func NewTLE(name string, line1 string, line2 string) TLE {
	return TLE{Name: name, Line1: line1, Line2: line2}
}

func ReadTLE(scanner *bufio.Scanner) (TLE, error) {
	// read the name
	tle := TLE{}
	if (scanner.Scan()) {
		tle.Name = strings.TrimSpace(scanner.Text())
	} else {
		return tle, io.EOF
	}

	if (scanner.Scan()) {
		tle.Line1 = strings.TrimSpace(scanner.Text())
	} else {
		return tle, errors.New("Line 1 missing")
	}

	if (scanner.Scan()) {
		tle.Line2 = strings.TrimSpace(scanner.Text())
	} else {
		return tle, errors.New("Line 2 missing")
	}
	return tle, nil
}

func ReadTLEs(r io.Reader) (map[string]TLE, error) {
	scanner := bufio.NewScanner(r)

	m := make(map[string]TLE)

	for {
		tle, err := ReadTLE(scanner)
		if err != nil {
			if err == io.EOF {
				return m, nil
			} else {
				return m, err
			}
		}

		m[tle.Name] = tle
	}

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return m, nil
}
