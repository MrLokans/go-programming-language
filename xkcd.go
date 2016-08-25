package main

import (
	"net/http"
	// "fmt"
	"encoding/json"
	"log"
	"os"
	"path"
	"path/filepath"
	"text/template"
)

const INDEX_FOLDER string = "index"
const XKCDUrl = "https://xkcd.com/571/info.0.json"

const xkcdTemplate = `
{{ .Title }} - {{ .Number }}
Url: {{ .Image }}

Transcript:
{{ .Transcript }}

Date: {{ .Day }}.{{ .Month }}.{{ .Year }}
`

// (['img', 'year', 'alt', 'num', 'safe_title', 'news', 'link', 'day', 'month', 'transcript', 'title'])

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

func main() {

	var absolutePath = path.Join(".", "index")
	absolutePath, err := filepath.Abs(absolutePath)
	if err != nil {
		log.Fatal("Error creating path")
		return
	}

	err = os.MkdirAll(absolutePath, os.ModeDir)
	if err != nil {
		log.Fatalf("Error creating dir %s, exiting", absolutePath)
	}

	log.Printf("Path created: %s", absolutePath)

	resp, err := http.Get(XKCDUrl)

	defer resp.Body.Close()

	if err != nil {
		log.Printf("Error occured downloading url %s", XKCDUrl)
		return
	}

	var xkcdOutput XKCDInfo

	err = json.NewDecoder(resp.Body).Decode(&xkcdOutput)

	if err != nil {
		log.Fatalf("Error decoding JSON for url %s: %v", XKCDUrl, err)
	}

	xkcd, err := template.New("xkcd").Parse(xkcdTemplate)
	if err != nil {
		log.Fatal(err)
	}

	if err := xkcd.Execute(os.Stdout, xkcdOutput); err != nil {
		log.Fatal(err)
	}

}
