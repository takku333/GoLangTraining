// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type omdbJSON struct {
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	write      string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Metascore  string
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string
	Response   string
	Error      string
}

func main() {
	if len(os.Args) == 1 {
		log.Fatal("set query")
	}
	name := os.Args[1]
	for _, word := range os.Args[2:] {
		name += " " + word
	}
	values := url.Values{}
	values.Add("t", name)
	reqURL := "https://omdbapi.com/?" + values.Encode()
	fmt.Println(reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		fmt.Println("HTTP request :" + err.Error())
	}

	var result omdbJSON
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		fmt.Println("json decode :" + err.Error())
	}
	resp.Body.Close()

	if result.Response == "False" {
		fmt.Println(result.Error)
		return
	}
	imgURL := result.Poster
	if imgURL == "N/A" {
		fmt.Println("No poster data")
		return
	}
	posterResp, err := http.Get(imgURL)
	if err != nil {
		fmt.Println("Poster request :" + err.Error())
	}
	poster, err := ioutil.ReadAll(posterResp.Body)
	posterResp.Body.Close()
	if err != nil {
		fmt.Println("Read poster data :" + err.Error())
	}
	filename := result.Title + ".jpg"
	ioutil.WriteFile(filename, poster, 0644)

}
