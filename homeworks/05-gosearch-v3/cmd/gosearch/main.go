package main

import (
	"flag"
	"fmt"
	"go-core-2/homeworks/05-gosearch-v3/pkg/crawler/spider"
	"go-core-2/homeworks/05-gosearch-v3/pkg/index"
	"go-core-2/homeworks/05-gosearch-v3/pkg/storage"
	"go-core-2/homeworks/05-gosearch-v3/pkg/storage/filestore"
	"log"
	"os"
)

const path = "filestore.txt"

type gosearch struct {
	scanner *spider.Service
	sites   []string
	depth   int
	index   *index.Index
	store   storage.Interface
}

func main() {
	var token = flag.String("s", "", "search for a particular word/token")
	flag.Parse()
	if *token == "" {
		fmt.Println("exiting as no token to search for was provided by input")
		return
	}

	s := new()
	f, err := os.Open(path)

	// вот здесь ниже как бы мне поэлегантее можно было написать?
	// ведь файла может и не быть... Проверить в if тип ошибки, а Retrieve вынести за if?
	if err == nil {
		docs, _ := s.store.Retrieve(f)
		s.index.Append(docs)
	}
	defer f.Close()

	fmt.Println("Processing...")

	if s.index.IsEmpty() {
		for _, url := range s.sites {
			od, err := s.scanner.Scan(url, s.depth)
			if err != nil {
				log.Println("error when scanning a site:", err)
			}
			s.index.Append(od)
		}
	}

	// build index for the documents
	s.index.Build()

	fmt.Println("Search results:")
	docs := s.index.Search(token)
	for _, d := range docs {
		fmt.Println(d)
	}

	// save the indexed docs into the file storage
	w, err := os.Create(path)
	if err != nil {
		log.Println("couldn't create a file to store results", err)
	}
	err = s.store.Save(w, s.index.All())
	if err != nil {
		log.Println("couldn't save results", err)
	}
	w.Close()
}

func new() *search {
	s := search{}
	s.sites = []string{"https://go.dev", "https://golang.org/"}
	s.depth = 2
	s.scanner = spider.New()
	s.index = index.New()
	s.store = filestore.New()
	return &s
}
