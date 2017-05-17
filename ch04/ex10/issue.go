// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	now := time.Now()
	var lessThanMonth []*github.Issue
	var lessThanYear []*github.Issue
	var moreThanYear []*github.Issue
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	monthBefor := now.AddDate(0, -1, 0)
	yearBefor := now.AddDate(-1, 0, 0)

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		if item.CreatedAt.Before(yearBefor) {
			moreThanYear = append(moreThanYear, item)
		} else if item.CreatedAt.Before(monthBefor) {
			lessThanYear = append(lessThanYear, item)
		} else {
			lessThanMonth = append(lessThanMonth, item)
		}
	}

	fmt.Println("before less than a month")
	for _, item := range lessThanMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("before less than a year")
	for _, item := range lessThanYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("before more than a year")
	for _, item := range lessThanMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
