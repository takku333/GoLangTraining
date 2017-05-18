// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("start")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func handler(w http.ResponseWriter, r *http.Request) {
	getGithunStatus(w)
}

func getGithunStatus(w http.ResponseWriter) {
	milestones, err := SearchMilestones(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "MilestoneList\n")
	for _, m := range milestones {
		fmt.Fprintf(w, "----------------------------------------\n")
		fmt.Fprintf(w, "Number : %d\n", m.Number)
		fmt.Fprintf(w, "Creator: %s\n", m.Creator.Login)
		fmt.Fprintf(w, "Title  : %s\n", m.Title)
		fmt.Fprintf(w, "Age    : %d\n", daysAgo(m.CreatedAt))
	}

	issues, err := SearchIssues(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "\nBug report\n")
	for _, i := range issues.Items {
		for _, label := range i.Labels {
			if label.Name == "bug" {
				fmt.Fprintf(w, "----------------------------------------\n")
				fmt.Fprintf(w, "Number: %d\n", i.Number)
				fmt.Fprintf(w, "User  : %s\n", i.User.Login)
				fmt.Fprintf(w, "Title : %s\n", i.Title)
				fmt.Fprintf(w, "Age   : %d\n", daysAgo(i.CreatedAt))
			}
		}
	}

	users, err := SearchUsers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "\nUsers\n")
	for _, y := range users {
		fmt.Fprintf(w, "----------------------------------------\n")
		fmt.Fprintf(w, "name  : %s\n", y.Login)
	}
}

func SearchMilestones(term string) ([]Milestone, error) {
	resp, err := http.Get("https://api.github.com/repos/" + term + "/milestones")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result []Milestone
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}

func SearchIssues(term string) (*IssuesSearchResult, error) {
	resp, err := http.Get(IssuesURL + "?q=" + term)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

func SearchUsers() ([]User, error) {
	resp, err := http.Get(UsersURL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result []User
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return result, nil
}
