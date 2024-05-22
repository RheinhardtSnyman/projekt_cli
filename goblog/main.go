package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/russross/blackfriday"
)

type Page struct {
	Title      string
	LastChange time.Time
	Content    string
}

type Pages []Page

var (
	flagSrcFolder  = flag.String("src", "./seiten/", "blog folder")
	flagTmplFolder = flag.String("tmpl", "./templates?", "template folder")
)

func main() {

}

func loadPage(fpath string) (Page, error) {
	var p Page
	fi, err := os.Stat(fpath)
	if err != nil {
		return p, fmt.Errorf("loadPage: %w", err)
	}
	p.Title = fi.Name()
	p.LastChange = fi.ModTime()
	b, err := os.ReadFile(fpath)
	if err != nil {
		return p, fmt.Errorf("loadPage.ReadFile: %w", err)
	}
	p.Content = string(blackfriday.MarkdownCommon(b))
	return p, nil
}

func loadPages(src string) (Pages, error) {
	var ps Pages
	fs, err := os.ReadDir(src)
	if err != nil {
		return ps, fmt.Errorf("LoadPages.ReadDir: %w", err)
	}
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		fpath := filepath.Join(src, f.Name())
		p, err := loadPage(fpath)
		if err != nil {
			return ps, fmt.Errorf("loadPages.loadPage: %w", err)
		}
		ps = append(ps, p)
	}
	return ps, nil
}
