// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"time"
)

const MilestoneURL = "https://api.github.com/repos/"
const IssuesURL = "https://api.github.com/search/issues"
const UsersURL = "https://api.github.com/users"

type Milestones struct {
	Milestones []*Milestone
}
type Milestone struct {
	Number    int
	URL       string
	Title     string
	Creator   *Actor
	CreatedAt time.Time `json:"created_at"`
}

type Actor struct {
	Login string
	URL   string
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	Labels    []*Labels
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Labels struct {
	Id   int
	Name string
}
