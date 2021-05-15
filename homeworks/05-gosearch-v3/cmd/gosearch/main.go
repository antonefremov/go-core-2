package main

import (
	"flag"
	"fmt"
	"go-core-2/homeworks/05-gosearch-v3/pkg/crawler/spider"
	"go-core-2/homeworks/05-gosearch-v3/pkg/index"
	"log"
)

type search struct {
	scanner *spider.Service
	sites   []string
	depth   int
	store   *index.Store
}

func main() {
	var token = flag.String("s", "", "search for a particular word/token")
	flag.Parse()
	if *token == "" {
		fmt.Println("exiting as no token to search for was provided by input")
		return
	}

	s := new()
	fmt.Println("Processing...")

	if s.store.IsEmpty() {
		for _, url := range s.sites {
			od, err := s.scanner.Scan(url, s.depth)
			if err != nil {
				log.Println("error when scanning a site:", err)
			}
			s.store.Append(od)
		}
	}

	// index the documents and sort them
	s.store.Index()
	s.store.Sort()

	fmt.Println("Search results:")
	docs := s.store.Search(token)
	for _, d := range docs {
		fmt.Println(d)
	}

	// save the existing docs for future use
	s.store.Save()
}

func new() *search {
	s := search{}
	s.sites = []string{"https://go.dev", "https://golang.org/"}
	s.depth = 2
	s.scanner = spider.New()
	s.store = index.New()
	return &s
}
