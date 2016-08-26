// The JSON-based web service of the Open Movie Database lets you search
// https://omdbapi.com/ for a movie by name and download its poster image. Write a tool
// 'poster' that downloads the poster image for the movie named on the command line.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	// "github.com/op/go-logging"
	"log"
	"net/http"
)

type MovieData struct {
	Title     string
	Year      string
	Rated     string
	Released  string
	Runtime   string
	Genre     string
	Director  string
	Writer    string
	Actors    string
	Plot      string
	Language  string
	Country   string
	Awards    string
	Poster    string
	Metascore int
	IMDBRate  string `json:"imdbRating"`
	IMDBID    string `json:"imdbID"`
	Type      string
	Response  string
}

const BASE_URL = "http://www.omdbapi.com"

// var log = logging.MustGetLogger("poster")

var movieQuery string

func init() {
	flag.StringVar(&movieQuery, "name", "", "Name of the movie to be searched")
}

func posterURL(title string) (string, error) {
	searchUrl := constructSearchURL(title)
	log.Printf("Examining URL %s", searchUrl)
	resp, err := http.Get(searchUrl)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	var result MovieData
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.Poster, nil
}

func constructSearchURL(title string) string {
	return BASE_URL + "/?t=" + title + "&r=json"
}

func main() {
	flag.Parse()
	if movieQuery == "" {
		fmt.Println("Emty movie query supplied, quitting")
		return
	}
	log.Printf("Movie %s poster is going to be searched", movieQuery)
	poster, err := posterURL(movieQuery)
	if err != nil {
		log.Fatalf("Error downloading poster %s: %v", movieQuery, err)
		return
	}
	log.Printf(poster)

}
