// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type xkcd struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func main() {
	if _, err := os.Stat("index"); err != nil {
		if err := os.Mkdir("index", 0777); err != nil {
			fmt.Println("mkdir :" + err.Error())
			os.Exit(1)
		}
		fmt.Println("mkdir \"index\"")
	}

	makeIndex(getMostRecentComicNum())

	args := os.Args
	if len(args) == 1 {
		log.Fatal("set query")
	}
	args = os.Args[1:]
	for _, query := range args {
		result := searchXKCD(query)
		fmt.Println("===============================")
		fmt.Printf("URL :https://xkcd.com/%d/info.0.json\ntranscript :%s\n", result.Num, result.Transcript)

	}
}

func makeIndex(num int) {
	fmt.Println("makeIndex")
	getch := make(chan int, 50)
	fileNum := 1
	for i := 1; i <= num; i++ {
		if jsonFileIsExisting(i) || i == 404 {
			fileNum++
			continue
		}
		url := "https://xkcd.com/" + strconv.Itoa(i) + "/info.0.json"
		go fetch(url, getch)
	}
	for fileNum < num {
		fileNum += <-getch
	}
	fmt.Println("makeIndex finish")
}

func fetch(url string, getch chan<- int) {
	result, err := getJSON(url)
	if err != nil {
		log.Fatal(err)
	}
	makeJsonfile(result)
	getch <- 1
}

func jsonFileIsExisting(i int) bool {
	path := "index/" + strconv.Itoa(i) + ".json"
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

func makeJsonfile(result *xkcd) {
	num := result.Num
	path := "index/" + strconv.Itoa(num) + ".json"
	jsonData, err := json.Marshal(result)
	if err = ioutil.WriteFile(path, jsonData, 0644); err != nil {
		fmt.Printf("write file :%v\n", err)
	}
	fmt.Println("make file :" + path)
}

func searchXKCD(index string) xkcd {
	jsonData, err := ioutil.ReadFile("index/" + index + ".json")
	if err != nil {
		fmt.Printf("read index : %v\n", err)
	}
	var result xkcd
	if err := json.Unmarshal(jsonData, &result); err != nil {
		fmt.Printf("Unmarshal :%v\n", err)
	}
	return result
}

func getJSON(url string) (*xkcd, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search %s faild: %s", url, resp.Status)
	}
	var result xkcd
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func getMostRecentComicNum() int {
	recenrtComicURL := "http://xkcd.com/info.0.json"
	recentComic, err := getJSON(recenrtComicURL)
	if err != nil {
		log.Fatalf("get recent comic json:%v", err)
	}
	return recentComic.Num
}
