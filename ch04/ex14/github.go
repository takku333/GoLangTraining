// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"time"
)

const MilestoneURL = "https://api.github.com/repos/"

type Milestones struct {
	Milestones  []*Milestone
}
type Milestone struct {
	Number    int
	URL   string 
	Title     string
	Creator   *Actor
	CreatedAt time.Time `json:"created_at"`
}

type Actor struct {
	Login   string
	URL string 
}