// Author: "Takumi Miyagawa"
// Copyright © 2017 RICOH Co, Ltd. All rights reserved

package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	listWithLoop := visit(nil, doc)
	sort.Strings(listWithLoop)

	listWithRecursive := recursiveVisit(nil, doc)
	sort.Strings(listWithRecursive)

	if !reflect.DeepEqual(listWithLoop, listWithRecursive) {
		fmt.Println("recursiveVisit is not correct!!!")
		for i := range listWithLoop {
			fmt.Printf("visit :%s\nrecursiveVisit :%s\n", listWithLoop[i], listWithRecursive[i])
		}
	}
	for _, s := range listWithRecursive {
		fmt.Println(s)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func recursiveVisit(links []string, n *html.Node) []string {
	if n != nil {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					links = append(links, a.Val)
				}
			}
		}
		links = recursiveVisit(links, n.FirstChild)
		links = recursiveVisit(links, n.NextSibling)
	}

	return links
}
