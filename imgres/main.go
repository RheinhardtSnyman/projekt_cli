package main

//.\imgres.exe -in ./imgres_bsp/ -size 400x400

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

var (
	flagInFolder  = flag.String("in", "./", "Input-Ordner")
	flagOutFolder = flag.String("out", "", "Output-Ordner")
	flagSize      = flag.String("size", "500x500", "maximale Größe")
)

func main() {
	flag.Parse()
	size, err := parseSize(*flagSize)
	if err != nil {
		fmt.Println("cannot create picSize: ", err)
		os.Exit(1)
	}
	outFolder := *flagSize
	if *flagOutFolder != "" {
		outFolder = *flagOutFolder
	}
	err = resizeFolderImages(*flagInFolder, outFolder, size)
	if err != nil {
		fmt.Println(err)
		os.Exit(10)
	}
}

type errorList struct {
	errs []error
}

func (e *errorList) add(err error) {
	if err != nil {
		e.errs = append(e.errs, err)
	}
}

func (e *errorList) hasErrors() bool {
	return len(e.errs) > 0
}

func (e *errorList) Error() string {
	if !e.hasErrors() {
		return ""
	}
	out := fmt.Sprintf("number of errors %d\n", len(e.errs))
	for i, err := range e.errs {
		out = fmt.Sprintf("%s\n%d: $s", out, i, err.Error())
	}
	return out
}

func resizeFolderImages(inFolder, outFolder string, size picSize) error {
	err := os.MkdirAll(outFolder, os.FileMode(0755))
	if err != nil {
		return fmt.Errorf("cannot create output folder: ", err)
	}
	dir, err := os.ReadDir(inFolder)
	if err != nil {
		return fmt.Errorf("cannot read from source: ", err)
	}
	errList := &errorList{}
	for _, fi := range dir {
		if fi.IsDir() || !useFile(fi.Name()) {
			continue
		}
		inPath := filepath.Join(inFolder, fi.Name())
		infile, err := os.Open(inPath)
		if err != nil {
			errList.add(fmt.Errorf("error opening file: %w", err))
			continue
		}
		outPath := filepath.Join(outFolder, fi.Name())
		outFile, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY, 0777)
		if err != nil {
			errList.add(fmt.Errorf("cannot create file: %w", err))
			infile.Close()
			continue
		}
		err = resize(size, infile, outFile)
		if err != nil {
			errList.add(fmt.Errorf("error resizing image: %w", err))
		}
		outFile.Close()
		infile.Close()
	}
	if errList.hasErrors() {
		return errList
	}
	return nil
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

func resize(ps picSize, r io.Reader, w io.Writer) error {
	img, format, err := image.Decode(r)
	if err != nil {
		return fmt.Errorf("error decoding: %w", err)
	}
	if format != "jpeg" {
		return fmt.Errorf("only jpeg is supported")
	}
	resized := imaging.Fit(
		img,
		ps.width, ps.height,
		imaging.Lanczos,
	)
	return jpeg.Encode(w, resized, nil)
}

func useFile(filename string) bool {
	allowed := []string{".jpg", ".jpeg"}
	ext := filepath.Ext(filename)
	for _, e := range allowed {
		if strings.EqualFold(ext, e) {
			return true
		}
	}
	return false
}
