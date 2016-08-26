package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	// "text/template"
)

const INDEX_FOLDER string = "index"
const XKCDUrl = "https://xkcd.com"

const xkcdTemplate = `
{{ .Title }} - {{ .Number }}
Url: {{ .Image }}

Transcript:
{{ .Transcript }}

Date: {{ .Day }}.{{ .Month }}.{{ .Year }}
`

type XKCDInfo struct {
	Title           string `json:"title"`
	Image           string `json:"img"`
	Year            string `json:"year"`
	AlternativeText string `json:"alt"`
	Number          int    `json:"number"`
	SafeTitle       string `json:"safe_title"`
	News            string `json:"news"`
	Day             string `json:"day"`
	Month           string `json:"month"`
	Transcript      string `json:"transcript"`
	Url             string `json:"link"`
}

func generateInfoFilePath(indexPath string, pageNumber int) string {
	var filename = fmt.Sprintf("%d", pageNumber) + ".json"
	return path.Join(indexPath, filename)
}

func generateXKCDInfoURL(XKCDNumber int) string {
	s := []string{XKCDUrl, fmt.Sprintf("%d", XKCDNumber), "info.0.json"}
	return strings.Join(s, "/")
}

func createIndexDir(indexPath string) error {
	err := os.MkdirAll(indexPath, os.ModeDir)
	if err != nil {
		log.Fatalf("Error creating dir %s, exiting", indexPath)
		return err
	}
	return nil
}

func downloadMetadata(indexPath string, pageNumber int, wg *sync.WaitGroup) error {
	url := generateXKCDInfoURL(pageNumber)
	outputFile := generateInfoFilePath(indexPath, pageNumber)
	defer wg.Done()
	log.Printf("Downloading url %s...", url)
	resp, err := http.Get(url)
	// defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Error downloading comic from %s: %v", url, err)
		return err
	}
	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error downloading comic from %s: %v", url, err)
		return err
	}
	defer out.Close()
	io.Copy(out, resp.Body)
	return nil
}

var comicCount int

func init() {
	flag.IntVar(&comicCount, "count", 10, "Number of XKCD comics to process.")
}

func main() {
	flag.Parse()

	indexPath, err := filepath.Abs(INDEX_FOLDER)
	if err != nil {
		log.Fatal("Error getting absolute path for %s", indexPath)
		return
	}

	err = createIndexDir(indexPath)
	if err != nil {
		log.Fatalf("Error creating dir %s, exiting", indexPath)
	}

	log.Printf("Path created: %s", indexPath)

	var wg sync.WaitGroup

	downloads := 0
	for downloads < comicCount {
		downloads++
		wg.Add(1)
		go downloadMetadata(indexPath, downloads, &wg)
	}
	wg.Wait()

}
