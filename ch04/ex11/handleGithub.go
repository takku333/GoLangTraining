// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"flag"
	"log"

	"context"

	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

const (
	owner = "takku333"
	repo  = "goplCh04Ex11"
	token = "51ef2f276adbfbcb32af5e74db6f57b89eefc54f"
)

var action = flag.String("action", "make", "select[make, read, update, close]")

func main() {
	flag.Parse()
	names := flag.Args()
	if len(names) != 1 {
		log.Fatal("set only one arg")
	}
	name := names[0]
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)

	switch *action {
	case "make":
		body := "gopl issue作成テスト用"
		opt := &github.IssueRequest{
			Title: &name,
			Body:  &body,
		}

		issue, _, err := client.Issues.Create(ctx, owner, repo, opt)
		if err != nil {
			log.Fatalf("create issue: %v", err)
		}
		fmt.Printf("create issue %q finish", *issue.Title)
	case "read":
		issues, _, err := client.Issues.ListByRepo(ctx, owner, repo, nil)
		if err != nil {
			log.Fatalf("read issues: %v", err)
		}
		for _, issue := range issues {
			if name == *issue.Title {
				fmt.Println(issue)
				return
			}
		}
		fmt.Println("no such issue")
		fmt.Println("[issue list]")
		for _, issue := range issues {
			fmt.Println(*issue.Title)
		}
	}

	// repos, _, err := client.Repositories.List("", nil)
	// if _, ok := err.(*github.RateLimitError); ok{
	// 	log.Fatal("hit rate limit")
	// }

}
