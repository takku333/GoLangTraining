// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"log"
	"os"
	"text/template"
	"time"
	"encoding/json"
	"fmt"
	"net/http"
	//"net/url"
	//"strings"
)

const templ = `----------------------------------------
Number: {{.Number}}
User:   {{.Creator.Login}}
Title:  {{.Title | printf "%.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

//!-daysAgo

//!+exec
var report = template.Must(template.New("milestonelist").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ))

func main() {
	result, err := SearchMilestones(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

func SearchMilestones(terms []string) (*Milestones, error) {
	//q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get("https://api.github.com/repos/golang/go/milestones")
	if err != nil {
		return nil, err
	}
	
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result Milestones
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}