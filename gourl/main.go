// go run .\main.go -header -o foo/bar/test.html http://go.dev
package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

var (
	flagOutput = flag.String("o", "", "Datei als Ziel")
	flagHeader = flag.Bool("header", false, "output of the HTTP-header")
)

func main() {
	flag.Parse()
	// fmt.Println(*flagOutput, os.Args, flag.Args())
	args := flag.Args()
	var w io.Writer
	// default ist stout
	w = os.Stdout
	if len(args) != 1 {
		fmt.Println("Bitte nur eine URL als Eingabe")
		os.Exit(1)
	}
	url := args[0]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Fehler beim Laden der Daten", url, err)
		os.Exit(2)
	}
	defer resp.Body.Close()
	if *flagOutput != "" {
		path := filepath.Dir(*flagOutput)
		err = os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("Ordnerpfad kann nicht erstellt werden", path, err)
			os.Exit(22)
		}
		//fd file descriptor
		fd, err := os.OpenFile(*flagOutput, os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			fmt.Println("Fehler beim Erstellen der Datei", err, *flagOutput)
			os.Exit(20)
		}
		defer fd.Close()
		w = fd
	}

	if *flagHeader {
		// fmt.Printf("%#v", resp.Header)
		for k, v := range resp.Header {
			fmt.Printf("%s:\n%v\n\n", k, v)
			// fmt.Printf("%s : %s", k, v)
		}
	}
	io.Copy(w, resp.Body)
}

func validateURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}
	// validURL := regexp.MustCompile("^http(s)?://[[:graph:]]+")
	// return validURL.MatchString(s)
	return true
}
