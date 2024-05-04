package main

import (
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		log.Println("Mindestens eine Datei als Parameter erwartet")
	}
	for _, fileName := range os.Args[1:] {
		fd, err := os.Open(fileName)
		if err != nil {
			log.Println("Fehler beim Öffnen der Datei: ", fileName, err)
			os.Exit(2)
		}
		_, err = io.Copy(os.Stdout, fd)
		if err != nil {
			log.Println("Error with io.Copy()", fileName)
		}
		fd.Close()
	}
}
