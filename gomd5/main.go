package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

var (
	flagFile = flag.String("file", "", "Wenn gesetzt wird direkt dieses File verwendet")
	flagURL  = flag.String("url", "", "Wenn gesetzt wird direkt HTTP get verwendet")
)

var input io.Reader

func main() {
	flag.Parse()
	switch {
	case *flagFile != "":
		f, err := os.Open(*flagFile)
		if err != nil {
			fmt.Println("Fehler bein Öffnen des Files: ", *flagFile, err)
			os.Exit(123)
		}
		defer f.Close()
		input = f
	case *flagURL != "":
		resp, err := http.Get(*flagURL)
		if err != nil {
			fmt.Println("Fehler beim URL: ", *flagURL, err)
			os.Exit(234)
		}
		defer resp.Body.Close()
		input = resp.Body
	default:
		input = os.Stdin
	}
	printMD5(input, os.Stdout)
}

func printMD5(r io.Reader, w io.Writer) {
	// b, _ := io.ReadAll(r)
	// fmt.Printf("%s", b)
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		fmt.Println("Fehler beim Einlesen: ", err)
		os.Exit(1)
	}
	fmt.Fprintf(w, "%x", h.Sum(nil))

}
