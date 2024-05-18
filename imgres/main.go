package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var (
	flagInFolder  = flag.String("in", "./", "Input-Ordner")
	flagOutFolder = flag.String("out", "", "Output-Ordner")
	flagSize      = flag.String("size", "500x500", "maximale Größe")
)

func main() {
	flag.Parse()
}

type picSize struct {
	width, height int
}

func parseSize(s string) (picSize, error) {
	var ps picSize
	parts := strings.Split(s, "x")
	if len(parts) != 2 {
		return ps, fmt.Errorf("%s does not fit to widthxheight", s)
	}
	var err error
	ps.width, err = strconv.Atoi(parts[0])
	if err != nil {
		return ps, fmt.Errorf("parseSize: ps.x: %w", err)
	}
	ps.height, err = strconv.Atoi(parts[1])
	if err != nil {
		return ps, fmt.Errorf("parseSize: ps.y: %w", err)
	}
	return ps, nil
}
